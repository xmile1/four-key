package helpers

//import "math"

func GetLeadTime(metricTags []tagMetricData) []tagMetricData {

	for i := 0; i < len(metricTags); i++ {
		var tagLeadTime float64 = 0
		if metricTags[i].featCommits != nil {
			for k := 0; k < len(metricTags[i].featCommits); k++ {
				tagLeadTime += metricTags[i].tagDate.Sub(metricTags[i].featCommits[k].Committer.When).Seconds()
			}
			var average = tagLeadTime / float64(len(metricTags[i].featCommits))
			metricTags[i].tagLeadTimeSeconds = average
		}
	}

	return metricTags
}

func GetMeanTimeToRestore(metricTags []tagMetricData) []tagMetricData {

	for i := len(metricTags) - 1; i >= 0; i-- {

		// tagMeanTimeRestoreAverageSeconds: How long does it take the longest fix to get to production after a release, because the tickets are worked on in series
		if metricTags[i].fixCommits != nil {
			var tagMeanTimeToRestore float64
			for k := 0; k < len(metricTags[i].fixCommits); k++ {
				time := metricTags[i].fixCommits[k].Committer.When.Sub(metricTags[i].tagDate).Seconds()
				if time > tagMeanTimeToRestore {
					tagMeanTimeToRestore = time
				}
			}
			metricTags[i].tagMeanTimeRestoreAverageSeconds = tagMeanTimeToRestore
		}
	}

	return metricTags
}

func GetChangeFailPercentage(metricTags []tagMetricData) []tagMetricData {

	for i := 0; i < len(metricTags); i++ {
		var totalFixCount = 0
		var totalFeatureCount = 0

		if metricTags[i].fixCommits != nil {
			totalFixCount += len(metricTags[i].fixCommits)
		}
		if metricTags[i].featCommits != nil {
			totalFeatureCount += len(metricTags[i].featCommits)
		}
		if totalFeatureCount != 0 {
			metricTags[i].tagChangeFailPercentage = float64(totalFixCount) / float64(totalFeatureCount) * 100
		}
	}

	return metricTags
}

func GetChangeFailPercentageByReleases(metricTags []tagMetricData) []tagMetricData {
	for i := len(metricTags) - 1; i >= 0; i-- {
		if len(metricTags[i].fixCommits) > 0 {
			metricTags[i].tagChangeFailPercentage = 100
		} else {
			metricTags[i].tagChangeFailPercentage = 0
		}
	}

	return metricTags
}
