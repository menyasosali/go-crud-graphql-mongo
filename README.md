# go-crud-graphql-mongo
CRUD project using MongoDB and Graphql

Get All Jobs
query GetAllJobs{ 
  jobs{ 
  _id 
  title 
  description 
  company 
  url 
 } 
}

=======================

Create Job
mutation CreateJobListing($input: CreateJobListingInput!){ 
   createJobListing(input:$input){ 
      _id 
      title 
      description 
      company 
      url 
   } 
 }

{ 
  "input": { 
    "title": "Software Development Engineer - I", 
    "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt", 
    "company": "Google", 
    "url": "www.google.com/"
  } 
}

=========================

Get Job By Id
query GetJob($id: ID!){ 
  job(id:$id){ 
    _id 
    title 
    description 
    url 
    company 
  } 
}

{ 
  "id": "638051d7acc418c13197fdf7" 
}

=========================

Update Job By Id
mutation UpdateJob($id: ID!,$input: UpdateJobListingInput!) { 
  updateJobListing(id:$id,input:$input){ 
    title 
    description 
    company 
    url 
 } 
}

{ 
  "id": "638051d3acc418c13197fdf6", 
  "input": { 
    "title": "Software Development Engineer - III" 
  } 
}

=================================

Delete Job By Id
mutation DeleteQuery($id: ID!) { 
  deleteJobListing(id:$id){ 
    deletedJobId 
  } 
}

{ 
  "id": "638051d3acc418c13197fdf6" 
}
