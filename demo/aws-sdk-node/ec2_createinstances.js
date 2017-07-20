var AWS = require('aws-sdk');
AWS.config.loadFromPath('./config.json');

var ec2 = new AWS.EC2({apiVersion: '2016-11-15'});
var params = {
  ImageId: 'ami-40290820',// 'ami-73f7da13', //ami-40290820
  InstanceType: 't1.micro',
  MinCount: 1,
  MaxCount: 1
};

ec2.runInstances(
  params,
  function(err, data) {
    if (err) {
      console.log('could not create instance', err);
      return;
    }
    var instanceId = data.Instances[0].InstanceId;
    console.log('Created instance', instanceId);

    params = {
      Resources: [instanceId],
      Tags: [{Key: 'Name', Value: 'SDK Sample'}]
    };
    ec2.createTags(
      params, function(err) {
        console.log('Tagging instance:', err ? 'failure' : 'success');
      }
    );
  }
)
