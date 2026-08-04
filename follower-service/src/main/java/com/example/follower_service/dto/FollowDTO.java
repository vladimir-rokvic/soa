package com.example.follower_service.dto;

public class FollowDTO {
    //Id client-a
    private String clientId;
    //Id korisnika kojeg client zeli da zaprati
    private String userId;

    public String getClientId() {
        return clientId;
    }

    public void setClientId(String clientId) {
        this.clientId = clientId;
    }

    public String getUserId() {
        return userId;
    }

    public void setUserId(String userId) {
        this.userId = userId;
    }
}
