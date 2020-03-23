package main

import "github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"

// searchTag looks for tagKey between tags
func searchTag(tags []*resourcegroupstaggingapi.Tag, tagKey *string) bool {
	for _, tag := range tags {
		if *tag.Key != *tagKey {
			return true
		}
	}
	return false
}
