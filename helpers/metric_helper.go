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

	for i := 0; i < len(metricTags); i++ {
		if metricTags[i].fixCommits != nil {
			var tagMeanTimeToRestore float64
			for k := 0; k < len(metricTags[i].fixCommits); k++ {
				tagMeanTimeToRestore += metricTags[i].fixCommits[k].Committer.When.Sub(metricTags[i].tagDate).Seconds()
			}
			var average = (tagMeanTimeToRestore) / float64(len(metricTags[i].fixCommits))
			metricTags[i].tagMeanTimeRestoreAverageSeconds = average
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

	var totalReleaseCount = 0
	var noOfReleasesWithFix = 0
	for i := 0; i < len(metricTags); i++ {

		if len(metricTags[i].fixCommits) > 0 {
			noOfReleasesWithFix = 1
		}
		if metricTags[i].releaseCommits != nil {
			totalReleaseCount += len(metricTags[i].releaseCommits)
		}

		if totalReleaseCount != 0 {
			metricTags[i].tagChangeFailPercentage = float64(noOfReleasesWithFix) * 100
			totalReleaseCount = 0
			noOfReleasesWithFix = 0
		}
	}

	return metricTags
}
