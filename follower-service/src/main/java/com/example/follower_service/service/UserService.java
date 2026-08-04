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

    public User followUser(FollowDTO dto) {
        return userRepository.followUser(dto.getClientId(), dto.getUserId());
    }

    public boolean doesHeFollow(String clientId, String userId) {
        return userRepository.isFollowing(clientId, userId);
    }
}
