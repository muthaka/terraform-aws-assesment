from flask import Flask
import boto3
instance_app = Flask(__name__)

region_name="us-west-2"
ec2 = boto3.client('ec2', region_name=region_name)

@instance_app.route('/tags')
def instance_tags():
    instance_tags = {}
    running_instances = ec2.describe_instances()
    for reservation in running_instances['Reservations']:
        for instance in reservation['Instances']:
            for tag in instance['Tags']:
                if tag['Key'] == 'Name':
                    instance_tags['Tag_Name'] = tag['Value']
                if tag['Key'] == 'Owner':
                    instance_tags['Tag_Owner'] = tag['Value']

    return instance_tags


@instance_app.route('/shutdown')
def instance_shutdown():
    tag_filters=[{'Name': 'tag:Name', 'Values': ['Flugel']}]
    instances_to_stop = []
    running_instances = ec2.describe_instances(Filters=tag_filters)
    for reservation in running_instances.get('Reservations'):
        for instance in reservation.get('Instances'):
            instances_to_stop.append(instance.get('InstanceId'))
    response = ec2.stop_instances(InstanceIds=instances_to_stop)
    return response


if __name__ == "__main__":
    instance_app.run(host="0.0.0.0", port=8000)