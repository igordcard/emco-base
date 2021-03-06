package module_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	dcm "gitlab.com/project-emco/core/emco-base/src/dcm/pkg/module"
	common "gitlab.com/project-emco/core/emco-base/src/orchestrator/common"
	"gitlab.com/project-emco/core/emco-base/src/orchestrator/pkg/infra/db"
	orch "gitlab.com/project-emco/core/emco-base/src/orchestrator/pkg/module"
	types "gitlab.com/project-emco/core/emco-base/src/orchestrator/pkg/module/types"
)

var _ = Describe("Cluster", func() {

	var (
		mdb    *db.NewMockDB
		client *dcm.ClusterClient
	)

	BeforeEach(func() {
		client = dcm.NewClusterClient()
		mdb = new(db.NewMockDB)
		mdb.Err = nil
		mdb.Items = []map[string]map[string][]byte{}
		db.DBconn = mdb
	})
	Describe("Cluster operations", func() {
		Context("from an empty database", func() {
			BeforeEach(func() {
				// create project in mocked db
				okey := orch.ProjectKey{
					ProjectName: "project",
				}
				p := orch.Project{}
				p.MetaData = orch.ProjectMetaData{
					Name:        "project",
					Description: "",
					UserData1:   "",
					UserData2:   "",
				}
				mdb.Insert("orchestrator", okey, nil, "projectmetadata", p)
				// create logical cloud in mocked db
				lkey := common.LogicalCloudKey{
					Project:          "project",
					LogicalCloudName: "logicalcloud",
				}
				lc := common.LogicalCloud{}
				lc.MetaData = types.Metadata{
					Name:        "logicalcloud",
					Description: "",
					UserData1:   "",
					UserData2:   "",
				}
				lc.Specification = common.Spec{
					NameSpace: "anything",
					Level:     "1",
				}
				mdb.Insert("orchestrator", lkey, nil, "logicalcloud", lc)
			})
			It("creation should succeed and return the resource created", func() {
				Skip("temporarily disabled")
				cluster := _createTestCluster("testcluster")
				cluster, err := client.CreateCluster("project", "logicalcloud", cluster)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(cluster.MetaData.Name).To(Equal("testcluster"))
				Expect(cluster.MetaData.Description).To(Equal(""))
				Expect(cluster.MetaData.UserData1).To(Equal(""))
				Expect(cluster.MetaData.UserData2).To(Equal(""))
			})
			// TODO
			// It("creation on instantiated cloud should fail", func() {
			// 	cluster := _createTestCluster("testcluster")
			// 	cluster, err := client.CreateCluster("project", "logicalcloud", cluster)
			// 	Expect(err).ShouldNot(HaveOccurred())
			// 	Expect(cluster.MetaData.Name).To(Equal("testcluster"))
			// 	Expect(cluster.MetaData.Description).To(Equal(""))
			// 	Expect(cluster.MetaData.UserData1).To(Equal(""))
			// 	Expect(cluster.MetaData.UserData2).To(Equal(""))
			// })
			It("get should fail and not return anything", func() {
				cluster, err := client.GetCluster("project", "logicalcloud", "testcluster")
				Expect(err).Should(HaveOccurred())
				Expect(cluster).To(Equal(common.Cluster{}))
			})
			It("create followed by get should return what was created", func() {
				Skip("temporarily disabled")
				cluster := _createTestCluster("testcluster")
				_, _ = client.CreateCluster("project", "logicalcloud", cluster)
				cluster, err := client.GetCluster("project", "logicalcloud", "testcluster")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(cluster).To(Equal(cluster))
			})
			It("create followed by get-all should return only what was created", func() {
				Skip("temporarily disabled")
				cluster := _createTestCluster("testcluster")
				_, _ = client.CreateCluster("project", "logicalcloud", cluster)
				clusters, err := client.GetAllClusters("project", "logicalcloud")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(len(clusters)).To(Equal(1))
				Expect(clusters[0]).To(Equal(cluster))
			})
			It("three creates followed by get-all should return all that was created", func() {
				Skip("temporarily disabled")
				cluster1 := _createTestCluster("testcluster1")
				cluster2 := _createTestCluster("testcluster2")
				cluster3 := _createTestCluster("testcluster3")
				_, _ = client.CreateCluster("project", "logicalcloud", cluster1)
				_, _ = client.CreateCluster("project", "logicalcloud", cluster2)
				_, _ = client.CreateCluster("project", "logicalcloud", cluster3)
				clusters, err := client.GetAllClusters("project", "logicalcloud")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(len(clusters)).To(Equal(3))
				Expect(clusters[0]).To(Equal(cluster1))
				Expect(clusters[1]).To(Equal(cluster2))
				Expect(clusters[2]).To(Equal(cluster3))
			})
			It("delete after creation should succeed and database remain empty", func() {
				Skip("temporarily disabled")
				cluster := _createTestCluster("testcluster")
				_, _ = client.CreateCluster("project", "logicalcloud", cluster)
				err := client.DeleteCluster("project", "logicalcloud", "testcluster")
				Expect(err).ShouldNot(HaveOccurred())
				clusters, err := client.GetAllClusters("project", "logicalcloud")
				Expect(len(clusters)).To(Equal(0))
			})
			It("delete when nothing exists should fail", func() {
				err := client.DeleteCluster("project", "logicalcloud", "testcluster")
				Expect(err).Should(HaveOccurred())
			})
			It("update after creation should succeed and return updated resource", func() {
				Skip("temporarily disabled")
				cluster := _createTestCluster("testcluster")
				_, _ = client.CreateCluster("project", "logicalcloud", cluster)
				cluster.MetaData.UserData1 = "new user data"
				cluster, err := client.UpdateCluster("project", "logicalcloud", "testcluster", cluster)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(cluster.MetaData.Name).To(Equal("testcluster"))
				Expect(cluster.MetaData.Description).To(Equal(""))
				Expect(cluster.MetaData.UserData1).To(Equal("new user data"))
				Expect(cluster.MetaData.UserData2).To(Equal(""))
			})
			It("create followed by updating the name is disallowed and should fail", func() {
				cluster := _createTestCluster("testcluster")
				_, _ = client.CreateCluster("project", "logicalcloud", cluster)
				cluster.MetaData.Name = "updated"
				cluster, err := client.UpdateCluster("project", "logicalcloud", "testcluster", cluster)
				Expect(err).Should(HaveOccurred())
				Expect(cluster).To(Equal(common.Cluster{}))
			})
		})
	})
})

// TODO:
// - test when cluster references already exist
// - appcontext status check for creation and deletion of cluster references
// - test GetClusterConfig

// _createTestCluster is an helper function to reduce code duplication
func _createTestCluster(name string) common.Cluster {
	return common.Cluster{
		MetaData: types.Metadata{
			Name:        name,
			Description: "",
			UserData1:   "",
			UserData2:   "",
		},
	}
}
