# AWS tagged resource finder
This small software allows you to search for tagged resources in a specified region.

You can specify the tag key you are looking for, the region and if you are looking for tagged or untagged resources.

## Usage examples:

Find resources with the tag Name in the eu-west-1 region
```bash
go run *.go -region eu-west-1 -tag Name
```

Find resources without the tag Name in the eu-west-1 region
```bash
go run *.go -region eu-west-1 -tag Name -untagged
```