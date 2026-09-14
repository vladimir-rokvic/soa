package com.example.follower_service.dto;

public class FollowDTO {
    //Id client-a
    private String clientId;
    //Id korisnika kojeg client zeli da zaprati
    private String userId;

    //ista fora kao i ranije
    private String clientUsername;
    private String userUsername;

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

    public String getClientUsername() {
        return clientUsername;
    }

    public void setClientUsername(String clientUsername) {
        this.clientUsername = clientUsername;
    }

    public String getUserUsername() {
        return userUsername;
    }

    public void setUserUsername(String userUsername) {
        this.userUsername = userUsername;
    }
}
