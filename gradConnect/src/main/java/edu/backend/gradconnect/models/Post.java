package edu.backend.gradconnect.models;

import jakarta.persistence.*;
import lombok.Data;

import java.util.List;
import java.util.Set;

@Entity
@Data
public class Post {

    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    Long id;
    String title;
    String description;
    String content;
    @Enumerated(EnumType.STRING)
    PostType type;
    @Enumerated(EnumType.STRING)
    PostAccess access;

    @ManyToOne
    @JoinColumn(name = "user_id")
    User user;

    @ManyToMany
    @JoinTable(
            name = "post_interest",
            joinColumns = @JoinColumn(name = "post_id"),
            inverseJoinColumns = @JoinColumn(name = "interest_id")
    )
    Set<Interest> interests;
    public Post() {}
}
