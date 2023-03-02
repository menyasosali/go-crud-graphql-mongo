package database

import (
	"context"
	"github.com/menyasosali/go-crud-graphql-mongo/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func (db *DB) GetJobs() []*model.JobListing {
	jobsCollect := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var jobListings []*model.JobListing

	filter := bson.D{}
	result, err := jobsCollect.Find(ctx, filter)
	if err != nil {
		log.Fatalf("error occurred while listing job items")
	}

	if err = result.All(context.TODO(), &jobListings); err != nil {
		panic(err)
	}

	return jobListings

}

func (db *DB) GetJob(jobId string) *model.JobListing {
	jobsCollect := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": _id}

	var jobListing model.JobListing

	err := jobsCollect.FindOne(ctx, filter).Decode(&jobListing)
	if err != nil {
		log.Fatalf("error occurred while listing jobs: %s", err)
	}

	return &jobListing
}

func (db *DB) CreateJobListing(jobInfo model.CreateJobListingInput) *model.JobListing {
	jobCollect := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	inserted, err := jobCollect.InsertOne(ctx, bson.M{
		"title":       jobInfo.Title,
		"description": jobInfo.Description,
		"company":     jobInfo.Company,
		"url":         jobInfo.URL,
	})
	if err != nil {
		log.Fatalf("error occurred while inserted job information: %s", err)
	}

	insertedID := inserted.InsertedID.(primitive.ObjectID).Hex()

	jobListing := model.JobListing{
		ID:          insertedID,
		Title:       jobInfo.Title,
		Description: jobInfo.Description,
		Company:     jobInfo.Company,
		URL:         jobInfo.URL,
	}

	return &jobListing
}

func (db *DB) UpdateJobListing(jobId string, jobInfo model.UpdateJobListingInput) *model.JobListing {
	jobCollect := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	updateJobInfo := bson.M{}

	if jobInfo.Title != nil {
		updateJobInfo["title"] = jobInfo.Title
	}

	if jobInfo.Description != nil {
		updateJobInfo["description"] = jobInfo.Description
	}

	if jobInfo.URL != nil {
		updateJobInfo["url"] = jobInfo.URL
	}

	_id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateJobInfo}

	results := jobCollect.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var jobListing model.JobListing

	if err := results.Decode(&jobListing); err != nil {
		log.Fatalf("error occurred while updating job information: %s", err)
	}

	return &jobListing

}

func (db *DB) DeleteJobListing(jobId string) *model.DeleteJobResponse {
	jobCollect := db.client.Database("graphql-job-board").Collection("jobs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(jobId)
	filter := bson.M{"_id": _id}

	_, err := jobCollect.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatalf("error occurred while deleting job item")
	}

	return &model.DeleteJobResponse{DeleteJobID: jobId}
}
