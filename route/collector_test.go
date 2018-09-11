package route_test

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/api/core/v1"
	ext "k8s.io/api/extensions/v1beta1"
	"k8s.io/client-go/kubernetes"

	"code.cloudfoundry.org/eirini"
	. "code.cloudfoundry.org/eirini/route"
	"code.cloudfoundry.org/eirini/route/routefakes"
	"k8s.io/client-go/kubernetes/fake"
)

var _ = Describe("Collector", func() {

	Context("Start collecting routes", func() {

		var (
			collector   *Collector
			fakeClient  kubernetes.Interface
			scheduler   *routefakes.FakeTaskScheduler
			workChannel chan []RegistryMessage
			routes      []string
			host        string
			serviceName string
		)

		const (
			appName      = "dora"
			namespace    = "testing"
			kubeEndpoint = "asgard"
			httpPort     = 80
			tlsPort      = 443
		)

		// handcraft json in order not to mirror the production implementation
		asJSONArray := func(uris []string) string {
			quotedUris := []string{}
			for _, uri := range uris {
				quotedUris = append(quotedUris, fmt.Sprintf("\"%s\"", uri))
			}

			return fmt.Sprintf("[%s]", strings.Join(quotedUris, ","))
		}

		createService := func(appName string) *v1.Service {
			service := &v1.Service{}

			service.Name = serviceName
			service.Namespace = namespace

			service.Annotations = map[string]string{
				"routes": asJSONArray(routes),
			}

			return service
		}

		createRule := func(serviceName string) ext.IngressRule {
			return ext.IngressRule{
				Host: host,
				IngressRuleValue: ext.IngressRuleValue{
					HTTP: &ext.HTTPIngressRuleValue{
						Paths: []ext.HTTPIngressPath{
							{
								Backend: ext.IngressBackend{
									ServiceName: serviceName,
								},
							},
						},
					},
				},
			}
		}

		createIngress := func(serviceNames ...string) *ext.Ingress {
			rules := []ext.IngressRule{}
			for _, name := range serviceNames {
				rule := createRule(name)
				rules = append(rules, rule)
			}
			return &ext.Ingress{
				Spec: ext.IngressSpec{
					Rules: rules,
				},
			}
		}

		createFakes := func() {
			service := createService(appName)
			_, err := fakeClient.CoreV1().Services(namespace).Create(service)
			Expect(err).ToNot(HaveOccurred())

			serviceNoAnnotation := &v1.Service{}
			serviceNoAnnotation.Name = "no-annotation"
			serviceNoAnnotation.Namespace = namespace

			_, err = fakeClient.CoreV1().Services(namespace).Create(serviceNoAnnotation)
			Expect(err).ToNot(HaveOccurred())

			ingress := createIngress(serviceName, "no-annotation")
			_, err = fakeClient.ExtensionsV1beta1().Ingresses(namespace).Create(ingress)
			Expect(err).ToNot(HaveOccurred())
		}

		BeforeEach(func() {
			serviceName = eirini.GetInternalServiceName(appName)
			host = "app.bosh.com"
			routes = []string{"route1.app.com", "route2.app.com"}

			scheduler = new(routefakes.FakeTaskScheduler)
			workChannel = make(chan []RegistryMessage, 1)
			fakeClient = fake.NewSimpleClientset()

			createFakes()
		})

		JustBeforeEach(func() {
			collector = &Collector{
				Client:        fakeClient,
				Scheduler:     scheduler,
				Work:          workChannel,
				KubeNamespace: namespace,
				KubeEndpoint:  kubeEndpoint,
			}

			collector.Start()
			task := scheduler.ScheduleArgsForCall(0)
			err := task()
			Expect(err).ToNot(HaveOccurred())
		})

		It("should use the scheduler to collect routes", func() {
			Expect(scheduler.ScheduleCallCount()).To(Equal(1))
		})

		It("should send the correct RegistryMessage in the work channel", func() {
			actualMessages := <-workChannel
			expectedMessages := []RegistryMessage{
				{
					Host:    kubeEndpoint,
					URIs:    routes,
					Port:    httpPort,
					TLSPort: tlsPort,
					App:     serviceName,
				},
			}
			Expect(actualMessages).To(Equal(expectedMessages))
		})
	})
})
