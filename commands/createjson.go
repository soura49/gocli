package commands

import (
	"os"

	"github.com/spf13/cobra"
)

/*
Usage for the command is create
Description is need to create a aws CFT template in json format
*/

func CreateJsonFile() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "To Create the AWS Cloudformation Json Template",
		RunE: func(cmd *cobra.Command, args []string) error {
			f, err := os.Create("create.json")
			defer f.Close()
			if err != nil {
				return err
			}
			d1 := `{
				"Outputs": {},
				"Description": "Jenkins installation with the AutoScaling Group",
				"AWSTemplateFormatVersion": "2010-09-09",
				"Resources": {
				"xxxxxJenkins": {
				"Type": "AWS::AutoScaling::AutoScalingGroup",
				"Properties": {
				"LaunchConfiguration": {
				"Ref": "xxxxxJenkinsLC"
				},
				"DesiredCapacity": "1",
				"MinSize": "1",
				"MaxSize": "3",
				"HealthCheckGracePeriod": "300",
				"HealthCheckType": "EC2",
				"NotificationTypes": [
				"autoscaling:EC2_INSTANCE_LAUNCH",
				"autoscaling:EC2_INSTANCE_LAUNCH_ERROR",
				"autoscaling:EC2_INSTANCE_TERMINATE",
				"autoscaling:EC2_INSTANCE_TERMINATE_ERROR"
				]
				},
				"Tags": [
				{
				"Key": "tr:application-asset-insight-id",
				"Value": "204392"
				},
				{
				"Key": "tr:enviroment-type",
				"Value": "Nonpord"
				},
				{
				"Key": "tr:finacial-identifier",
				"Value": "xxxx"
				},
				{
				"Key": "Name",
				"Value": "Jenkins_11282017"
				}
				],
				"VPCZoneIdentifier": [
				"Subnet-xxxxx"
				]
				}
				},
				"ScaleUpPolicy": {
				"Type": "AWS::AutoScaling::ScalingPolicy",
				"Properties": {
				"AdjustmentType": "ChangeIncapacity",
				"AutoScalingGroupName": {
				"Ref": "xxxxxJenkins"
				},
				"Cooldown": "360",
				"ScalingAdjustment": "1"
				}
				},
				"ScalingDownPolicy": {
				"Type": "AWS::AutoScaling::ScalingPolicy",
				"Properties": {
				"AdjustmentType": "ChangeIncapacity",
				"AutoScalingGroupName": {
				"Ref": "xxxxxJenkins"
				},
				"Cooldown": "360",
				"ScalingAdjustment": "-1"
				}
				},
				"CPUUtilizationHigh": {
				"Type": "AWS::CloudWatch::Alarm",
				"Properties": {
				"AlarmDescription": "cpu utlization high for auto scaling group triggering",
				"MetricName": "CpuUtilization",
				"NameSpace": "AWS/EC2",
				"Statistic": "Average",
				"period": "60",
				"EvaluationPeriods": "5",
				"Threshold": "80",
				"AlarmActions": [
				{
				"Ref": "ScaleUpPolicy"
				}
				],
				"Dimensions": [
				{
				"Name": "AutoScalingGroupName",
				"Value": {
				"Ref": "xxxxxJenkins"
				}
				}
				],
				"ComparsionOperator": "GreaterThanEqualToThreshold",
				"TreatMissingData": "missing"
				}
				},
				"CPUUtilizationDown": {
				"Type": "AWS::CloudWatch::Alarm",
				"Properties": {
				"AlarmDescription": "cpu utlization high for auto scaling group triggering",
				"MetricName": "CpuUtilization",
				"NameSpace": "AWS/EC2",
				"Statistic": "Average",
				"period": "60",
				"EvaluationPeriods": "5",
				"Threshold": "50",
				"AlarmActions": [
				{
				"Ref": "ScaleUpPolicy"
				}
				],
				"Dimensions": [
				{
				"Name": "AutoScalingGroupName",
				"Value": {
				"Ref": "xxxxxJenkins"
				}
				}
				],
				"ComparsionOperator": "LessThanEqualToThreshold",
				"TreatMissingData": "missing"
				}
				},
				"xxxxxJenkinsLC": {
				"Type": "AWS::AutoScaling:LaunchConfiguration",
				"Properties": {
				"AssociatePublicIpaddress": "True",
				"IamInstacnceProfile": "xxxxx-iamrole",
				"ImageId": "ami-dfadsf",
				"SecurityGroups": [
				"sg-xxxxx",
				"sg-xxxxxx"
				],
				"UserData": {
				"Fn::Base64": {
				"Fn::Join": [
				"\n",
				[
				"#!/bin/bash -V \n",
				"yum -y install update \n",
				"yum -y install java"
				]
				]
				}
				},
				"InstanceType": "t2.micro"
				}
			}
			}`
			f.WriteString(d1)
			return nil
		},
	}
}
