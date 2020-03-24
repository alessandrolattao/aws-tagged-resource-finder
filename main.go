package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
)

func main() {
	var err error

	// get the region flag
	regionFlag := flag.String("region", "", "AWS Region code")

	// get the tag key flag
	tagKeyFlag := flag.String("tag", "", "Tag key to find")

	// get the untagged flag
	untaggedFlag := flag.Bool("untagged", false, "Display resources without tag if set")

	// parse all the flags
	flag.Parse()

	// exit if missing parameters
	if *regionFlag == "" || *tagKeyFlag == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// create the session
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// create the ResourceGroupsTaggingAPI client
	client := resourcegroupstaggingapi.New(session, aws.NewConfig().WithRegion(*regionFlag))

	// pagination and Input/Output variables
	var paginationToken string = ""
	var resourcesInput *resourcegroupstaggingapi.GetResourcesInput
	var resourcesOutput *resourcegroupstaggingapi.GetResourcesOutput
	var counter int = 0

	// loop over resources
	for {

		// request input
		resourcesInput = &resourcegroupstaggingapi.GetResourcesInput{
			ResourcesPerPage: aws.Int64(100),
			PaginationToken:  &paginationToken,
		}

		// retrieve all resources
		resourcesOutput, err = client.GetResources(resourcesInput)
		if err != nil {
			fmt.Println(err)
		}

		// for every resource
		for _, resource := range resourcesOutput.ResourceTagMappingList {

			// analyze tags
			tagFound := tagExists(resource.Tags, tagKeyFlag)

			// if the resource is tagged and I want untagged resource then skip
			if tagFound && !*untaggedFlag {
				continue
			}

			// if the resource is not tagged and I want tagged resource then skip
			if !tagFound && *untaggedFlag {
				continue
			}

			printResource(resource)
			counter = counter + 1
		}

		// loop until the paginationToken is empty, no more pages
		paginationToken = *resourcesOutput.PaginationToken
		if *resourcesOutput.PaginationToken == "" {
			break
		}
	}

	fmt.Printf(`Looking for tag "%s" between resources in "%s"... `, *tagKeyFlag, *regionFlag)
	if *untaggedFlag {
		fmt.Printf("found %d untagged resources!\n", counter)
	} else {
		fmt.Printf("found %d tagged resources!\n", counter)
	}

}
