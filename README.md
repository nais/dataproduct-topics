# Dataproduct Topics
## Prerequisites
```
gcloud iam --project <project_id> \
  service-accounts create dataproduct-topics \
  --description="Manually created service-account for dataproduct-topics"

bq mk \
  --project_id <project_id> \
  --location europe-north1 \
  dataproduct_topics.dataproduct_topics \
  collection_time:datetime,pool:string,topic:string,team:string

bq add-iam-policy-binding \
  --member='serviceAccount:dataproduct-topics@<project_id>.iam.gserviceaccount.com' \
  --role='roles/bigquery.dataEditor' \
  <project_id>:dataproduct_topics.dataproduct_topics
```
