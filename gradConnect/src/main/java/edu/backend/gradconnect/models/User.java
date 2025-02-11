package edu.backend.gradconnect.models;

import jakarta.persistence.*;
import lombok.Data;

import java.util.List;
import java.util.Set;

@Entity
@Data
public class User {
    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private String firstName;

    private String lastName;

    @Column(unique = true)
    private String email;

    private String password;

    private String bio;

    private String profilePhotoUrl;

    @Enumerated(EnumType.STRING)
    private UserType userType;

    private Boolean isVerified;

    private Long connectedCount;

    @OneToMany(cascade=CascadeType.ALL)
    @JoinColumn(name = "posted_by")
    private List<Podcast> podcasts;

    @ManyToMany
    @JoinTable(
            name = "user_interest",
            joinColumns = @JoinColumn(name = "user_id"),
            inverseJoinColumns = @JoinColumn(name = "interest_id")
    )
    private Set<Interest> interests;

    @OneToMany(cascade = CascadeType.ALL, mappedBy = "user")
    private List<Post> posts;
    @OneToMany(cascade = CascadeType.ALL, mappedBy = "user")
    private List<JobPost> jobPosts;

    public User() {}
}

