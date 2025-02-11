package edu.backend.gradconnect.models;

import jakarta.persistence.*;
import lombok.Data;

import java.time.LocalDateTime;

@Entity
@Data
public class JobPost {
    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    Long id;

    @ManyToOne
    @JoinColumn(name = "user_id")
    User user;

    @Enumerated(EnumType.STRING)
    JobPostType jobPostType;
    String title;
    String role;
    String skills;
    String imageUrl;
    String description;
    String registrationLink;
    LocalDateTime closedAt;
}
