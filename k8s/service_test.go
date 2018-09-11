package k8s_test

import (
	"code.cloudfoundry.org/eirini"
	"code.cloudfoundry.org/eirini/models/cf"
	"code.cloudfoundry.org/eirini/opi"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"

	. "code.cloudfoundry.org/eirini/k8s"
)

var _ = Describe("Service", func() {

	var (
		fakeClient     kubernetes.Interface
		serviceManager ServiceManager
	)

	const (
		namespace = "midgard"
	)

	BeforeEach(func() {
		fakeClient = fake.NewSimpleClientset()
		serviceManager = NewServiceManager(fakeClient, namespace)
	})

	Context("When exposing an existing LRP", func() {

		var (
			lrp *opi.LRP
			err error
		)

		BeforeEach(func() {
			lrp = createLRP("baldur", "54321.0", `["my.example.route"]`)
		})

		Context("When creating a usual service", func() {

			JustBeforeEach(func() {
				err = serviceManager.Create(lrp)
			})

			It("should not fail", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should create a service", func() {
				serviceName := eirini.GetInternalServiceName("baldur")
				service, getErr := fakeClient.CoreV1().Services(namespace).Get(serviceName, meta.GetOptions{})
				Expect(getErr).ToNot(HaveOccurred())
				Expect(service).To(Equal(toService(lrp, namespace)))
			})

			Context("When recreating a existing service", func() {

				BeforeEach(func() {
					lrp = createLRP("baldur", "54321.0", `["my.example.route"]`)
				})

				JustBeforeEach(func() {
					err = serviceManager.Create(lrp)
				})

				It("should return an error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})

		Context("When creating a headless service", func() {

			JustBeforeEach(func() {
				err = serviceManager.CreateHeadless(lrp)
			})

			It("should not fail", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should create a service", func() {
				serviceName := eirini.GetInternalHeadlessServiceName("baldur")
				service, getErr := fakeClient.CoreV1().Services(namespace).Get(serviceName, meta.GetOptions{})
				Expect(getErr).ToNot(HaveOccurred())
				Expect(service).To(Equal(toHeadlessService(lrp, namespace)))
			})

			Context("When recreating a existing service", func() {

				BeforeEach(func() {
					lrp = createLRP("baldur", "54321.0", `["my.example.route"]`)
				})

				JustBeforeEach(func() {
					err = serviceManager.CreateHeadless(lrp)
				})

				It("should return an error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})

	Context("When deleting", func() {

		var service *v1.Service

		assertServiceIsDeleted := func(err error) {
			Expect(err).ToNot(HaveOccurred())

			services, err := fakeClient.CoreV1().Services(namespace).List(meta.ListOptions{})
			Expect(err).ToNot(HaveOccurred())
			Expect(services.Items).To(BeEmpty())
		}

		Context("a regular service", func() {

			var (
				err error
			)

			BeforeEach(func() {
				lrp := createLRP("odin", "1234.5", `["my.example.route"]`)
				service = toService(lrp, namespace)
				_, err = fakeClient.CoreV1().Services(namespace).Create(service)
				Expect(err).ToNot(HaveOccurred())
			})

			JustBeforeEach(func() {
				err = serviceManager.Delete("odin")
			})

			It("flags the service with delete", func() {
				service, err = fakeClient.CoreV1().Services(namespace).Get(service.Name, meta.GetOptions{})
				Expect(err).ToNot(HaveOccurred())

				Expect(service.Annotations["delete"]).To(Equal("true"))
			})

			It("moves the registered routes to unregistered", func() {
				service, err = fakeClient.CoreV1().Services(namespace).Get(service.Name, meta.GetOptions{})
				Expect(err).ToNot(HaveOccurred())

				Expect(service.Annotations[eirini.RegisteredRoutes]).To(Equal(`[]`))
				Expect(service.Annotations[eirini.UnregisteredRoutes]).To(Equal(`["my.example.route"]`))
			})

			Context("when the service does not exist", func() {

				JustBeforeEach(func() {
					err = serviceManager.Delete("tyr")
				})

				It("returns an error", func() {
					Expect(err).To(HaveOccurred())
				})
			})

			Context("When there are unregistered routes", func() {
				BeforeEach(func() {
					lrp := createLRP("odin", "1234.5", `["my-new.example.route"]`)
					err = serviceManager.Update(lrp)
					Expect(err).ToNot(HaveOccurred())
				})

				It("append the routes to the unregistered routes", func() {
					service, err = fakeClient.CoreV1().Services(namespace).Get(service.Name, meta.GetOptions{})
					Expect(err).ToNot(HaveOccurred())

					Expect(service.Annotations[eirini.RegisteredRoutes]).To(Equal(`[]`))
					Expect(service.Annotations[eirini.UnregisteredRoutes]).To(Equal(`["my.example.route","my-new.example.route"]`))

				})

			})
		})

		Context("a headless service", func() {

			var err error

			BeforeEach(func() {
				lrp := createLRP("odin", "1234.5", `["my.example.route"]`)
				service = toHeadlessService(lrp, namespace)
				_, err = fakeClient.CoreV1().Services(namespace).Create(service)
				Expect(err).ToNot(HaveOccurred())
			})

			JustBeforeEach(func() {
				err = serviceManager.DeleteHeadless("odin")
			})

			It("deletes the service", func() {
				assertServiceIsDeleted(err)
			})

			Context("when the service does not exist", func() {

				JustBeforeEach(func() {
					err = serviceManager.DeleteHeadless("tyr")
				})

				It("returns an error", func() {
					Expect(err).To(HaveOccurred())
				})

			})
		})
	})

	Context("When updating an service", func() {
		var (
			err            error
			lrp            *opi.LRP
			serviceName    string
			updatedService *v1.Service
		)

		BeforeEach(func() {
			lrp = createLRP("odin", "1234.5", `["my.example.route"]`)
			err = serviceManager.Create(lrp)
		})

		Context("when routes are updated", func() {

			JustBeforeEach(func() {
				err = serviceManager.Update(lrp)
				Expect(err).ToNot(HaveOccurred())

				serviceName = eirini.GetInternalServiceName("odin")

				updatedService, err = fakeClient.CoreV1().Services(namespace).Get(serviceName, meta.GetOptions{})
			})

			Context("When a route is replaced", func() {
				BeforeEach(func() {
					lrp = createLRP("odin", "1234.5", `["my-new.example.route"]`)
				})

				It("should update the routes annotation", func() {
					Expect(err).ToNot(HaveOccurred())
					Expect(updatedService.Annotations[eirini.RegisteredRoutes]).To(Equal(`["my-new.example.route"]`))
				})

				It("should remove the difference and add it to unregisteredRoutes annotation ", func() {
					Expect(err).ToNot(HaveOccurred())
					Expect(updatedService.Annotations[eirini.UnregisteredRoutes]).To(Equal(`["my.example.route"]`))
				})
			})

			Context("When routes are added", func() {
				BeforeEach(func() {
					lrp = createLRP("odin", "1234.5", `["my.example.route","my-new.example.route"]`)
				})

				It("should contain the old route", func() {
					Expect(err).ToNot(HaveOccurred())
					Expect(updatedService.Annotations[eirini.RegisteredRoutes]).To(ContainSubstring(`"my.example.route"`))
				})

				It("should contain the new route", func() {
					Expect(err).ToNot(HaveOccurred())
					Expect(updatedService.Annotations[eirini.RegisteredRoutes]).To(ContainSubstring(`"my-new.example.route"`))
				})

				It("should be empty", func() {
					Expect(err).ToNot(HaveOccurred())
					Expect(updatedService.Annotations[eirini.UnregisteredRoutes]).To(Equal(`[]`))
				})
			})

			Context("When routes are completly removed", func() {
				BeforeEach(func() {
					lrp = createLRP("odin", "1234.5", `[]`)
				})

				It("should empty the routes annotation", func() {
					Expect(err).ToNot(HaveOccurred())
					Expect(updatedService.Annotations[eirini.RegisteredRoutes]).To(Equal(`[]`))
				})

				It("should unregister the existing routes", func() {
					Expect(err).ToNot(HaveOccurred())
					Expect(updatedService.Annotations[eirini.UnregisteredRoutes]).To(ContainSubstring(`my.example.route`))
				})
			})
		})
	})

	Context("ListRoutes", func() {

		var (
			routes         []*eirini.Routes
			err            error
			updatedService *v1.Service
		)

		JustBeforeEach(func() {
			routes, err = serviceManager.ListRoutes()
		})

		Context("When there are existing services", func() {

			var lrp *opi.LRP

			BeforeEach(func() {
				lrp = createLRP("baldur", "54321.0", `["my.example.route"]`)
				err = serviceManager.Create(lrp)
				Expect(err).ToNot(HaveOccurred())
				lrp = createLRP("baldur", "54322.0", `["my-new.example.route"]`)
				err = serviceManager.Update(lrp)
				Expect(err).ToNot(HaveOccurred())
			})

			It("should not return an error", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should return the correct routes", func() {
				Expect(routes).To(HaveLen(1))
				route := routes[0]
				Expect(route.Routes).To(ContainElement("my-new.example.route"))
				Expect(route.UnregisteredRoutes).To(ContainElement("my.example.route"))
				Expect(route.Name).To(Equal(eirini.GetInternalServiceName("baldur")))
			})

			Context("When a route was unregistered", func() {
				JustBeforeEach(func() {
					route := routes[0]
					err = route.PopUnregisteredRoutes()
					Expect(err).ToNot(HaveOccurred())
				})

				It("should set the correct remove unregisterred route callback", func() {
					updatedService, err = fakeClient.CoreV1().Services(namespace).Get(eirini.GetInternalServiceName("baldur"), meta.GetOptions{})
					Expect(err).ToNot(HaveOccurred())
					Expect(updatedService.Annotations[eirini.UnregisteredRoutes]).To(Equal(`[]`))
				})

				Context("When the service is scheduled for deletion", func() {
					BeforeEach(func() {
						err = serviceManager.Delete("baldur")
						Expect(err).ToNot(HaveOccurred())
					})

					It("should delete the service", func() {
						_, err = fakeClient.CoreV1().Services(namespace).Get(eirini.GetInternalServiceName("baldur"), meta.GetOptions{})
						Expect(err).To(HaveOccurred())
					})
				})

			})

			Context("When there are headless services", func() {
				BeforeEach(func() {
					err = serviceManager.CreateHeadless(lrp)
					Expect(err).ToNot(HaveOccurred())
				})

				It("should not return an error", func() {
					Expect(err).ToNot(HaveOccurred())
				})

				It("should return only one Routes object", func() {
					Expect(routes).To(HaveLen(1))
				})
			})

			Context("When there are non cf services", func() {
				BeforeEach(func() {
					service := &v1.Service{}
					service.Name = "some-other-service"
					_, err = fakeClient.CoreV1().Services(namespace).Create(service)
					Expect(err).ToNot(HaveOccurred())
				})

				It("should not return an error", func() {
					Expect(err).ToNot(HaveOccurred())
				})

				It("should return only one Routes object", func() {
					Expect(routes).To(HaveLen(1))
				})
			})
		})
	})
})

func getServicesNames(services *v1.ServiceList) []string {
	serviceNames := []string{}
	for _, s := range services.Items {
		serviceNames = append(serviceNames, s.Name)
	}
	return serviceNames
}

func toService(lrp *opi.LRP, namespace string) *v1.Service {
	service := &v1.Service{
		Spec: v1.ServiceSpec{
			Ports: []v1.ServicePort{
				{
					Name: "service",
					Port: 8080,
				},
			},
			Selector: map[string]string{
				"name": lrp.Name,
			},
		},
	}

	service.Name = eirini.GetInternalServiceName(lrp.Name)
	service.Namespace = namespace
	service.Labels = map[string]string{
		"name": lrp.Name,
	}

	service.Annotations = map[string]string{
		eirini.RegisteredRoutes:   lrp.Metadata[cf.VcapAppUris],
		eirini.UnregisteredRoutes: `[]`,
	}

	return service
}

func toHeadlessService(lrp *opi.LRP, namespace string) *v1.Service {
	service := &v1.Service{
		Spec: v1.ServiceSpec{
			ClusterIP: "None",
			Ports: []v1.ServicePort{
				{
					Name: "service",
					Port: 8080,
				},
			},
			Selector: map[string]string{
				"name": lrp.Name,
			},
		},
	}

	service.Name = eirini.GetInternalHeadlessServiceName(lrp.Name)
	service.Namespace = namespace
	service.Labels = map[string]string{
		"name": lrp.Name,
	}

	return service
}
