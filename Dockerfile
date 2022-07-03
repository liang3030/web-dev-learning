FROM golang:1.18-alpine3.16 as build

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 go build -o ./mainApp ./.



FROM alpine:3.16

WORKDIR /app

COPY --from=build /app/mainApp /app/mainApp
COPY --from=build /app/index.html /app/index.html

ENV port=8080
ENV twilioFromNumber=""
ENV twilioAccountSid=""
ENV twilioAuthToken=""


CMD ./mainApp -port=${port} -twilioAuthToken=${twilioAuthToken} -twilioAccountSid=${twilioAccountSid} -twilioFromNumber=${twilioFromNumber}





