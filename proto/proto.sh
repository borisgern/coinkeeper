#!/bin/bash

protoc proto/expenses.proto --go_out=plugins=grpc:.