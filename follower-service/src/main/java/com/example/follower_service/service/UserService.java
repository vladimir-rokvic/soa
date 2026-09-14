package com.example.follower_service.service;

import com.example.follower_service.dto.FollowDTO;
import com.example.follower_service.model.User;
import com.example.follower_service.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class UserService {

    @Autowired
    private UserRepository userRepository;

    public List<User> getAllThatFollows(String userId) {
        return userRepository.findAllThatFollows(userId);
    }

    public User createUser(User user) {
        User nu = new User();
        if(user.getId() != null
        || !user.getId().isEmpty()) {
            nu.setId(user.getId());
        }
        nu.setUsername(user.getUsername());

        return userRepository.save(user);
    }

    //obicna pomocna funkcija nemoj se previse baviti ovime
    private boolean userExists(String id) {
        return userRepository.userExists(id);
    }

    public User followUser(FollowDTO dto) {
        //ako ne postoji kreiraj ga
        if(!userExists(dto.getClientId())) {
            User u = new User(dto.getClientId(), dto.getClientUsername());
            createUser(u);
        }

        if(!userExists(dto.getUserId())) {
            User u = new User(dto.getUserId(), dto.getUserUsername());
            createUser(u);
        }
        return userRepository.followUser(dto.getClientId(), dto.getUserId());
    }

    public boolean doesHeFollow(String clientId, String userId) {
        return userRepository.isFollowing(clientId, userId);
    }

    public List<User> getRecommendations(String userId) {
        List<User> users = userRepository.findRecommendations(userId);
        if(users.isEmpty()) {
            users = userRepository.findRecommendationsByCount(userId);
        }

        return users;
    }
}
