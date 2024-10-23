from googleapiclient import discovery
from oauth2client.client import GoogleCredentials

def create_instance(project, zone, name):
    credentials = GoogleCredentials.get_application_default()
    service = discovery.build('compute', 'v1', credentials=credentials)

    config = {
        'name': name,
        'machineType': f"zones/{zone}/machineTypes/e2-medium",
        'disks': [
            {
                'boot': True,
                'autoDelete': True,
                'initializeParams': {
                    'sourceImage': 'projects/debian-cloud/global/images/family/debian-10'
                }
            }
        ],
        'networkInterfaces': [
            {
                'network': 'global/networks/default',
                'accessConfigs': [
                    {'type': 'ONE_TO_ONE_NAT', 'name': 'External NAT'}
                ]
            }
        ]
    }

    return service.instances().insert(
        project=project,
        zone=zone,
        body=config).execute()

# project = '<your-gcp-project-id>'
project = 'loyal-rookery-439515-j4'
zone = 'us-central1-a'
name = 'api-instance'

create_instance(project, zone, name)