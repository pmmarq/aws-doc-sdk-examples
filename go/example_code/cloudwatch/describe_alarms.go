/*
   Copyright 2010-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.

   This file is licensed under the Apache License, Version 2.0 (the "License").
   You may not use this file except in compliance with the License. A copy of
   the License is located at

    http://aws.amazon.com/apache2.0/

   This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
   CONDITIONS OF ANY KIND, either express or implied. See the License for the
   specific language governing permissions and limitations under the License.
*/

sess, err := session.NewSession()
if err != nil {
    fmt.Println("failed to create session,", err)
    return
}

svc := cloudwatch.New(sess)

params := &cloudwatch.DescribeAlarmsInput{
    ActionPrefix:    aws.String("ActionPrefix"),
    AlarmNamePrefix: aws.String("AlarmNamePrefix"),
    AlarmNames: []*string{
        aws.String("AlarmName"), // Required
        // More values...
    },
    MaxRecords: aws.Int64(1),
    NextToken:  aws.String("NextToken"),
    StateValue: aws.String("ALARM"),
}
resp, err := svc.DescribeAlarms(params)

if err != nil {
    // Print the error, cast err to awserr.Error to get the Code and
    // Message from an error.
    fmt.Println(err.Error())
    return
}

//TODO: show token handling?
// Pretty-print the response data.
fmt.Println(resp)
