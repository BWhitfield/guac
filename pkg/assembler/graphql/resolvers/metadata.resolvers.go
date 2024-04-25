package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// IngestHasMetadata is the resolver for the ingestHasMetadata field.
func (r *mutationResolver) IngestHasMetadata(ctx context.Context, subject model.PackageSourceOrArtifactInput, pkgMatchType model.MatchFlags, hasMetadata model.HasMetadataInputSpec) (string, error) {
	funcName := "IngestHasMetadata"
	if err := validatePackageSourceOrArtifactInput(&subject, funcName); err != nil {
		return "", gqlerror.Errorf("%v ::  %s", funcName, err)
	}

	return r.Backend.IngestHasMetadata(ctx, subject, &pkgMatchType, hasMetadata)
}

// IngestBulkHasMetadata is the resolver for the ingestBulkHasMetadata field.
func (r *mutationResolver) IngestBulkHasMetadata(ctx context.Context, subjects model.PackageSourceOrArtifactInputs, pkgMatchType model.MatchFlags, hasMetadataList []*model.HasMetadataInputSpec) ([]string, error) {
	funcName := "IngestBulkHasMetadata"
	valuesDefined := 0
	if len(subjects.Packages) > 0 {
		if len(subjects.Packages) != len(hasMetadataList) {
			return []string{}, gqlerror.Errorf("%v :: uneven packages and hasMetadata for ingestion", funcName)
		}
		valuesDefined = valuesDefined + 1
	}
	if len(subjects.Artifacts) > 0 {
		if len(subjects.Artifacts) != len(hasMetadataList) {
			return []string{}, gqlerror.Errorf("%v :: uneven artifacts and hasMetadata for ingestion", funcName)
		}
		valuesDefined = valuesDefined + 1
	}
	if len(subjects.Sources) > 0 {
		if len(subjects.Sources) != len(hasMetadataList) {
			return []string{}, gqlerror.Errorf("%v :: uneven sources and hasMetadata for ingestion", funcName)
		}
		valuesDefined = valuesDefined + 1
	}
	if valuesDefined != 1 {
		return []string{}, gqlerror.Errorf("%v :: must specify at most packages, artifacts or sources", funcName)
	}
	return r.Backend.IngestBulkHasMetadata(ctx, subjects, &pkgMatchType, hasMetadataList)
}

// HasMetadata is the resolver for the HasMetadata field.
func (r *queryResolver) HasMetadata(ctx context.Context, hasMetadataSpec model.HasMetadataSpec) ([]*model.HasMetadata, error) {
	if err := validatePackageSourceOrArtifactQueryFilter(hasMetadataSpec.Subject); err != nil {
		return nil, gqlerror.Errorf("HasMetadata ::  %s", err)
	}
	return r.Backend.HasMetadata(ctx, &hasMetadataSpec)
}

// HasMetadataList is the resolver for the HasMetadataList field.
func (r *queryResolver) HasMetadataList(ctx context.Context, hasMetadataSpec model.HasMetadataSpec, after *string, first *int) (*model.HasMetadataConnection, error) {
	if err := validatePackageSourceOrArtifactQueryFilter(hasMetadataSpec.Subject); err != nil {
		return nil, gqlerror.Errorf("HasMetadata ::  %s", err)
	}
	return r.Backend.HasMetadataList(ctx, hasMetadataSpec, after, first)
}
