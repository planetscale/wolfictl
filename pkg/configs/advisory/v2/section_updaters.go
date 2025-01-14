package v2

import "github.com/wolfi-dev/wolfictl/pkg/configs"

func NewAdvisoriesSectionUpdater(
	updater configs.SectionUpdater[Advisories, Document],
) configs.EntryUpdater[Document] {
	yamlASTMutater := configs.NewTargetedYAMLASTMutater[Advisories, Document](
		"advisories",
		updater,
		func(doc Document, data Advisories) Document {
			doc.Advisories = data

			// Since we're using _this_ version of wolfictl to update the document, we
			// should update the schema version, which ensures that any features of the
			// current schema being used in this document update are accounted for in the
			// document's expressed schema version.
			doc.SchemaVersion = SchemaVersion

			return doc
		},
	)

	return configs.NewYAMLUpdateFunc[Document](yamlASTMutater)
}
