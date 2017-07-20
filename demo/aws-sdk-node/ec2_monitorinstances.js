var AWS = require('aws-sdk');
AWS.config.loadFromPath('./config.json');

ec2 = new AWS.EC2({apiVersion: '2016-11-15'});

var params = {
  InstanceIds: ['i-0431e6a97e03653a4'],
  DryRun: true
};

if (process.argv[2].toUpperCase() === 'ON') {
  ec2.monitorInstances(
    params,
    function(err, data) {
      if (err && err.code === 'DryRunOperation') {
        params.DryRun = false
        ec2.monitorInstances(
          params,
          function(err, data) {
            if (err) {
              console.log('Error', err);
            } else if (data) {
              console.log('Success', data.InstanceMonitorings);
            }
          }
        );
      } else {
        console.log('You don\'t have permission to change instance monitoring.');
      }
    }
  );
} else if (process.argv[2].toUpperCase() === 'OFF') {
  ec2.unmonitorInstances(
    params,
    function(err, data) {
      if (!(err && err.code === 'DryRunOperation')) {
        return console.log('You don\'t have permission to change instance monitoring.');
      }
      params.DryRun = false;
      ec2.unmonitorInstances(
        params,
        function(err, data) {
          if (err) {
            console.error('Error', err);
          } else if (data) {
            console.log('Success', data.InstanceMonitorings);
          }
        }
      )
    }
  )
}
