# AWS tagged resource finder ![Release](https://github.com/alessandrolattao/aws-tagged-resource-finder/workflows/Release/badge.svg?branch=master)
This small software allows you to search for tagged resources in a specified region.

You can specify the tag key you are looking for, the region and if you are looking for tagged or untagged resources.

This software needs awscli to be configured first.

## Usage examples:

Find resources with the tag Name in the eu-west-1 region
```bash
aws-tagged-resource-finder -region eu-west-1 -tag Name
```

Find resources without the tag Name in the eu-west-1 region
```bash
aws-tagged-resource-finder -region eu-west-1 -tag Name -untagged
```

## Download:

[https://github.com/alessandrolattao/aws-tagged-resource-finder/releases](https://github.com/alessandrolattao/aws-tagged-resource-finder/releases)