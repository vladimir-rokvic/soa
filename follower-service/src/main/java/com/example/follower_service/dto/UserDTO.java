package com.example.follower_service.dto;

import com.example.follower_service.model.User;

public class UserDTO {
    private String id;
    private String username;

    public UserDTO() {}

    public UserDTO(User u) {
        id = u.getId();
        username = u.getUsername();
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }
}
