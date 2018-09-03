package isolated

import (
	"code.cloudfoundry.org/cli/integration/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("auth command", func() {
	Context("Help", func() {
		It("displays the help information", func() {
			session := helpers.CF("auth", "--help")
			Eventually(session).Should(Say("NAME:"))
			Eventually(session).Should(Say("auth - Authenticate non-interactively\n\n"))

			Eventually(session).Should(Say("USAGE:"))
			Eventually(session).Should(Say("cf auth USERNAME PASSWORD\n"))
			Eventually(session).Should(Say("cf auth CLIENT_ID CLIENT_SECRET --client-credentials\n\n"))

			Eventually(session).Should(Say("ENVIRONMENT VARIABLES:"))
			Eventually(session).Should(Say("CF_USERNAME=user\\s+Authenticating user. Overridden if USERNAME argument is provided."))
			Eventually(session).Should(Say("CF_PASSWORD=password\\s+Password associated with user. Overriden if PASSWORD argument is provided."))

			Eventually(session).Should(Say("WARNING:"))
			Eventually(session).Should(Say("Providing your password as a command line option is highly discouraged"))
			Eventually(session).Should(Say("Your password may be visible to others and may be recorded in your shell history\n"))
			Eventually(session).Should(Say("Consider using the CF_PASSWORD environment variable instead\n\n"))

			Eventually(session).Should(Say("EXAMPLES:"))
			Eventually(session).Should(Say("cf auth name@example\\.com \"my password\" \\(use quotes for passwords with a space\\)"))
			Eventually(session).Should(Say("cf auth name@example\\.com \\\"\\\\\"password\\\\\"\\\" \\(escape quotes if used in password\\)\n\n"))

			Eventually(session).Should(Say("OPTIONS:"))
			Eventually(session).Should(Say("--client-credentials\\s+Use \\(non-user\\) service account \\(also called client credentials\\)\n\n"))

			Eventually(session).Should(Say("SEE ALSO:"))
			Eventually(session).Should(Say("api, login, target"))

			Eventually(session).Should(Exit(0))
		})
	})

	Context("when no positional arguments are provided", func() {
		Context("and no env variables are provided", func() {
			It("errors-out with the help information", func() {
				session := helpers.CF("auth")
				Eventually(session.Err).Should(Say("Username and password not provided."))
				Eventually(session).Should(Say("NAME:"))

				Eventually(session).Should(Exit(1))
			})
		})

		Context("when env variables are provided", func() {
			It("authenticates the user", func() {
				username, password := helpers.GetCredentials()
				env := map[string]string{
					"CF_USERNAME": username,
					"CF_PASSWORD": password,
				}
				session := helpers.CFWithEnv(env, "auth")

				Eventually(session).Should(Say("API endpoint: %s", helpers.GetAPI()))
				Eventually(session).Should(Say("Authenticating\\.\\.\\."))
				Eventually(session).Should(Say("OK"))
				Eventually(session).Should(Say("Use 'cf target' to view or set your target org and space"))

				Eventually(session).Should(Exit(0))
			})
		})
	})

	Context("when only a username is provided", func() {
		It("errors-out with a password required error and the help information", func() {
			session := helpers.CF("auth", "some-user")
			Eventually(session.Err).Should(Say("Password not provided."))
			Eventually(session).Should(Say("NAME:"))

			Eventually(session).Should(Exit(1))
		})
	})

	Context("when only a password is provided", func() {
		It("errors-out with a username required error and the help information", func() {
			env := map[string]string{
				"CF_PASSWORD": "some-pass",
			}
			session := helpers.CFWithEnv(env, "auth")
			Eventually(session.Err).Should(Say("Username not provided."))
			Eventually(session).Should(Say("NAME:"))

			Eventually(session).Should(Exit(1))
		})
	})

	Context("when too many arguments are provided", func() {
		It("displays an 'unknown flag' error message", func() {
			session := helpers.CF("auth", "some-username", "some-password", "-a", "api.bosh-lite.com")

			Eventually(session.Err).Should(Say("Incorrect Usage: unknown flag `a'"))
			Eventually(session).Should(Say("NAME:"))

			Eventually(session).Should(Exit(1))
		})
	})

	Context("when the API endpoint is not set", func() {
		BeforeEach(func() {
			helpers.UnsetAPI()
		})

		It("displays an error message", func() {
			session := helpers.CF("auth", "some-username", "some-password")

			Eventually(session).Should(Say("FAILED"))
			Eventually(session.Err).Should(Say("No API endpoint set\\. Use 'cf login' or 'cf api' to target an endpoint\\."))

			Eventually(session).Should(Exit(1))
		})
	})

	Context("when no flags are set (logging in with password grant type)", func() {
		Context("when the user provides an invalid username/password combo", func() {
			BeforeEach(func() {
				helpers.LoginCF()
				helpers.TargetOrgAndSpace(ReadOnlyOrg, ReadOnlySpace)
			})

			It("clears the cached tokens and target info, then displays an error message", func() {
				session := helpers.CF("auth", "some-username", "some-password")

				Eventually(session).Should(Say("API endpoint: %s", helpers.GetAPI()))
				Eventually(session).Should(Say("Authenticating\\.\\.\\."))
				Eventually(session).Should(Say("FAILED"))
				Eventually(session.Err).Should(Say("Credentials were rejected, please try again\\."))
				Eventually(session).Should(Exit(1))

				// Verify that the user is not logged-in
				targetSession1 := helpers.CF("target")
				Eventually(targetSession1.Err).Should(Say("Not logged in\\. Use 'cf login' to log in\\."))
				Eventually(targetSession1).Should(Say("FAILED"))
				Eventually(targetSession1).Should(Exit(1))

				// Verify that neither org nor space is targeted
				helpers.LoginCF()
				targetSession2 := helpers.CF("target")
				Eventually(targetSession2).Should(Say("No org or space targeted, use 'cf target -o ORG -s SPACE'"))
				Eventually(targetSession2).Should(Exit(0))
			})
		})

		Context("when the username and password are valid", func() {
			It("authenticates the user", func() {
				username, password := helpers.GetCredentials()
				session := helpers.CF("auth", username, password)

				Eventually(session).Should(Say("API endpoint: %s", helpers.GetAPI()))
				Eventually(session).Should(Say("Authenticating\\.\\.\\."))
				Eventually(session).Should(Say("OK"))
				Eventually(session).Should(Say("Use 'cf target' to view or set your target org and space"))

				Eventually(session).Should(Exit(0))
			})
		})
	})

	Context("when the 'client-credentials' flag is set", func() {
		Context("when the user provides an invalid client id/secret combo", func() {
			BeforeEach(func() {
				helpers.LoginCF()
				helpers.TargetOrgAndSpace(ReadOnlyOrg, ReadOnlySpace)
			})

			It("clears the cached tokens and target info, then displays an error message", func() {
				session := helpers.CF("auth", "some-client-id", "some-client-secret", "--client-credentials")

				Eventually(session).Should(Say("API endpoint: %s", helpers.GetAPI()))
				Eventually(session).Should(Say("Authenticating\\.\\.\\."))
				Eventually(session).Should(Say("FAILED"))
				Eventually(session.Err).Should(Say("Credentials were rejected, please try again\\."))
				Eventually(session).Should(Exit(1))

				// Verify that the user is not logged-in
				targetSession1 := helpers.CF("target")
				Eventually(targetSession1.Err).Should(Say("Not logged in\\. Use 'cf login' to log in\\."))
				Eventually(targetSession1).Should(Say("FAILED"))
				Eventually(targetSession1).Should(Exit(1))

				// Verify that neither org nor space is targeted
				helpers.LoginCF()
				targetSession2 := helpers.CF("target")
				Eventually(targetSession2).Should(Say("No org or space targeted, use 'cf target -o ORG -s SPACE'"))
				Eventually(targetSession2).Should(Exit(0))
			})
		})

		Context("when the client id and client secret are valid", func() {
			It("authenticates the user", func() {
				clientID, clientSecret := helpers.SkipIfClientCredentialsNotSet()
				session := helpers.CF("auth", clientID, clientSecret, "--client-credentials")

				Eventually(session).Should(Say("API endpoint: %s", helpers.GetAPI()))
				Eventually(session).Should(Say("Authenticating\\.\\.\\."))
				Eventually(session).Should(Say("OK"))
				Eventually(session).Should(Say("Use 'cf target' to view or set your target org and space"))

				Eventually(session).Should(Exit(0))
			})
		})
	})

	Context("when a user authenticates with valid client credentials", func() {
		BeforeEach(func() {
			clientID, clientSecret := helpers.SkipIfClientCredentialsNotSet()
			session := helpers.CF("auth", clientID, clientSecret, "--client-credentials")
			Eventually(session).Should(Exit(0))
		})

		Context("when a different user authenticates with valid password credentials", func() {
			It("should fail authentication and display an error informing the user they need to log out", func() {
				username, password := helpers.GetCredentials()
				session := helpers.CF("auth", username, password)

				Eventually(session).Should(Say("FAILED"))
				Eventually(session.Err).Should(Say("Service account currently logged in\\. Use 'cf logout' to log out service account and try again\\."))
				Eventually(session).Should(Exit(1))
			})
		})
	})
})
