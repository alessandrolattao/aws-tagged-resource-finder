package main

import "github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"

// tagExists looks for tagKey between tags
func tagExists(tags []*resourcegroupstaggingapi.Tag, tagKey *string) bool {
	for _, tag := range tags {
		if *tag.Key == *tagKey {
			return true
		}
	}
	return false
}
