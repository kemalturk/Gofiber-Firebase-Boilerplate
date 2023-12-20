deploy:
	gcloud config set project <ProjectID>
	gcloud functions deploy MyCloudFunction --runtime go120 --trigger-http