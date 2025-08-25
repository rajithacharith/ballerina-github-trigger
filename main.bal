import ballerinax/trigger.github;
import ballerina/log;

public configurable string webhookSecret = ?;

listener github:Listener github = new (listenerConfig = {webhookSecret: webhookSecret}, listenOn = 8000);

service github:PullRequestService on github {
    remote function onOpened(github:PullRequestEvent payload) returns error|() {
        log:printInfo("onOpened triggered");
        do {
        } on fail error err {
            // handle error
            return error("unhandled error", err);
        }
    }

    remote function onClosed(github:PullRequestEvent payload) returns error|() {
    log:printInfo("onClosed triggered");
    do {
        } on fail error err {
            // handle error
            return error("unhandled error", err);
        }
    }

    remote function onReopened(github:PullRequestEvent payload) returns error|() {
    log:printInfo("onReopened triggered");
    do {
        } on fail error err {
            // handle error
            return error("unhandled error", err);
        }
    }

    remote function onAssigned(github:PullRequestEvent payload) returns error|() {
    log:printInfo("onAssigned triggered");
    do {
        } on fail error err {
            // handle error
            return error("unhandled error", err);
        }
    }

    remote function onUnassigned(github:PullRequestEvent payload) returns error|() {
    log:printInfo("onUnassigned triggered");
    do {
        } on fail error err {
            // handle error
            return error("unhandled error", err);
        }
    }

    remote function onReviewRequested(github:PullRequestEvent payload) returns error|() {
    log:printInfo("onReviewRequested triggered");
    do {
        } on fail error err {
            // handle error
            return error("unhandled error", err);
        }
    }

    remote function onReviewRequestRemoved(github:PullRequestEvent payload) returns error|() {
    log:printInfo("onReviewRequestRemoved triggered");
    do {
        } on fail error err {
            // handle error
            return error("unhandled error", err);
        }
    }

    remote function onLabeled(github:PullRequestEvent payload) returns error|() {
    log:printInfo("onLabeled triggered");
    do {
        } on fail error err {
            // handle error
            return error("unhandled error", err);
        }
    }

    remote function onUnlabeled(github:PullRequestEvent payload) returns error|() {
    log:printInfo("onUnlabeled triggered");
    do {
        } on fail error err {
            // handle error
            return error("unhandled error", err);
        }
    }

    remote function onEdited(github:PullRequestEvent payload) returns error|() {
    log:printInfo("onEdited triggered");
    do {
        } on fail error err {
            // handle error
            return error("unhandled error", err);
        }
    }
}
