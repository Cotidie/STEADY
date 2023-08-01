# Google Cloud Platform
## Overview
![gcp](https://i.imgur.com/jTQaUbt.jpg)  
Google Cloud Platform is a suite of Computing, Networking, Storage, Machine Learning and Management tools which renders computing resources to developers with elaborated features. GCP has over 100 services spanning IaaS, PaaS and SaaS. GCP comes with GCP Shell which is a powerful tool for managing GCP services, running on a GCE VM resource.

### Hierachy
```plain
|- Organization
    |- Folder
        |- Project
            - Resource (Engine)
            - Resource (Storage)
        |- Project
            - Resource
        ...
```
GCP has a hierachy of Organization, Folder, Project and Resources.
- Organization is the top entity for all the project (e.g. company)
- Folder is a logical group of projects (e.g. team)
- Project is a billable unit of resources.
- Resource is an entity created by a GCP service. (e.g. storage, instance)

## QnA
**How to set up an organization?**  
: Go to \[IAM & Admin\] > \[Identity & Organization\] and follow the steps. It will require custom domain that has alive DNS.  