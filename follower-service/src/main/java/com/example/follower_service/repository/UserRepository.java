package com.example.follower_service.repository;

import com.example.follower_service.model.User;
import org.springframework.data.neo4j.repository.Neo4jRepository;
import org.springframework.data.neo4j.repository.query.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.util.List;

@Repository
public interface UserRepository extends Neo4jRepository<User, String> {
    @Query("MATCH (:User {id: $id}) -[:FOLLOWS]-> (u:User) RETURN u")
    List<User> findAllThatFollows(@Param("id") String id);

    @Query("""
           MATCH (client:User {id: $clientId}), (user:User {id: $userId})
           MERGE (client) -[:FOLLOWS]-> (user)
           RETURN client
           """)
    User followUser(@Param("clientId") String clientId, @Param("userId") String userId);

    @Query("""
           MATCH (client:User {id: $clientId}), (target:User {id: $userId})
           RETURN EXISTS { MATCH (client) -[:FOLLOWS]-> (target)} AS isFollowing
           """
    )
    boolean isFollowing(@Param("clientId") String clientId, @Param("userId") String userId);
}
