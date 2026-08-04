package com.example.follower_service.controller;

import com.example.follower_service.dto.FollowDTO;
import com.example.follower_service.model.User;
import com.example.follower_service.service.UserService;
import org.neo4j.cypherdsl.core.Use;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/followers/users")
public class UserController {

    @Autowired
    private UserService userService;

    //sve one koje korisnik prati
    @GetMapping("{id}/follows")
    public ResponseEntity<?> getAllThatFollows(@PathVariable("id") String id) {
        List<User> ret = userService.getAllThatFollows(id);

        return ResponseEntity.ok(ret);
    }

    @GetMapping("{clientId}/follows/{userId}")
    public ResponseEntity<?> doesHeFollow(
            @PathVariable("clientId") String clientId,
            @PathVariable("userId") String userId
    ) {
        return ResponseEntity.ok(userService.doesHeFollow(clientId, userId));
    }

    @PostMapping("/")
    public ResponseEntity<?> createUser(@RequestBody User user) {
        User u = userService.createUser(user);

        if(u == null) return ResponseEntity.badRequest().body("What");

        return ResponseEntity.ok(u);
    }

    @PostMapping("/follow")
    public ResponseEntity<?> followUser(@RequestBody FollowDTO dto) {
        User ret = userService.followUser(dto);

        if(ret == null) return ResponseEntity.badRequest().build();

        return ResponseEntity.ok(ret);
    }
}
