## Breaking Changes

### Removed Funcs

1. *CatalogBillingPlansItem.UnmarshalJSON([]byte) error
1. CatalogBillingPlansItem.MarshalJSON() ([]byte,error)

## Struct Changes

### Removed Structs

1. CatalogBillingPlansItem

## Signature Changes

### Struct Fields

1. Catalog.BillingPlans changed type from *[]CatalogBillingPlansItem to map[string][]ReservationBillingPlan

## New Content

### New Funcs

1. Catalog.MarshalJSON() ([]byte,error)
