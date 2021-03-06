package recipe_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"code.cloudfoundry.org/eirini"
	. "code.cloudfoundry.org/eirini/recipe"
)

var _ = Describe("Unzip function", func() {

	var (
		targetDir string
		srcZip    string
		err       error
		extractor eirini.Extractor
	)

	JustBeforeEach(func() {
		extractor = &Unzipper{}
		err = extractor.Extract(srcZip, targetDir)
	})

	Context("Unzip succeeds", func() {

		fileContents := map[string]string{
			"file1":                       "this is the content of test file 1",
			"innerDir/file2":              "this is the content of test file 2",
			"innerDir/innermostDir/file3": "this is the content of test file 3",
		}

		filePermissions := map[string]os.FileMode{
			"file1":                       0742,
			"innerDir/file2":              0651,
			"innerDir/innermostDir/file3": 0777,
		}

		getRoot := func(path string) string {
			pathParts := strings.Split(path, "/")
			return pathParts[0]
		}

		removeFile := func(file string) {
			ioErr := os.RemoveAll(file)
			Expect(ioErr).ToNot(HaveOccurred())
		}

		cleanUpFiles := func() {
			for filePath := range fileContents {
				rootDir := getRoot(filePath)
				removeFile(rootDir)
			}
		}

		assertFileContents := func(file string, expectedContent string) {
			path := filepath.Join(targetDir, file)
			content, ioErr := ioutil.ReadFile(path)
			Expect(ioErr).ToNot(HaveOccurred())
			Expect(content).To(Equal([]byte(expectedContent)))
		}

		assertFilePermissions := func(file string, expectedPermissions os.FileMode) {
			path := filepath.Join(targetDir, file)
			fileInfo, ioErr := os.Stat(path)
			Expect(ioErr).ToNot(HaveOccurred())
			Expect(fileInfo.Mode()).To(Equal(expectedPermissions))
		}

		assertFilesUnzippedSuccessfully := func() {
			It("should not fail", func() {
				Expect(err).ToNot(HaveOccurred())
			})

			It("should unzip the files in the target directory", func() {
				for fileName := range fileContents {
					path := filepath.Join(targetDir, fileName)
					Expect(path).To(BeAnExistingFile())
				}
			})

			It("should not change file contents", func() {
				for file, expectedContent := range fileContents {
					assertFileContents(file, expectedContent)
				}
			})

			It("should not change file permissions", func() {
				for file, expectedPermissions := range filePermissions {
					assertFilePermissions(file, expectedPermissions)
				}
			})
		}

		Context("When target directory is current working directory", func() {

			BeforeEach(func() {
				srcZip = "testdata/unzip_me.zip"
				targetDir = "."
			})

			AfterEach(func() {
				cleanUpFiles()
			})

			assertFilesUnzippedSuccessfully()
		})

		Context("When target directory is not empty string", func() {
			BeforeEach(func() {
				targetDir = "testdata/tmp"

				ioErr := os.Mkdir(targetDir, 0750)
				Expect(ioErr).ToNot(HaveOccurred())
			})

			AfterEach(func() {
				ioErr := os.RemoveAll(targetDir)
				Expect(ioErr).ToNot(HaveOccurred())
			})

			Context("When the zip does not contain the directory files", func() {
				BeforeEach(func() {
					srcZip = "testdata/just_files.zip"
				})
				assertFilesUnzippedSuccessfully()
			})

			Context("When the zip contains the directory files", func() {
				BeforeEach(func() {
					srcZip = "testdata/unzip_me.zip"
				})
				assertFilesUnzippedSuccessfully()
			})
		})
	})

	Context("Unzip fails", func() {

		Context("When target directory is not specified", func() {

			BeforeEach(func() {
				targetDir = ""
				srcZip = "testdata/unzip_me.zip"
			})

			It("should fail", func() {
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(ContainSubstring("target directory cannot be empty")))
			})
		})

		Context("When target dir is not a directory", func() {

			BeforeEach(func() {
				targetDir = "testdata/unzip_me.zip"
				srcZip = "testdata/unzip_me.zip"
			})

			It("should fail", func() {
				Expect(err).To(HaveOccurred())
			})

		})

		Context("When source zip archive does not exist", func() {

			BeforeEach(func() {
				targetDir = "testdata"
				srcZip = "non-existent"
			})

			It("should fail", func() {
				Expect(err).To(HaveOccurred())
			})

		})

		Context("When source is not a zip archive", func() {

			BeforeEach(func() {
				targetDir = "testdata"
				srcZip = "testdata/file.notzip"
			})

			It("should fail", func() {
				Expect(err).To(HaveOccurred())
			})

		})
	})

})
