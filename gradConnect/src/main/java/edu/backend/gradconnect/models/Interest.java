package edu.backend.gradconnect.models;

import jakarta.persistence.*;
import lombok.Data;
import java.util.Set;

@Entity
@Data
public class Interest {
    @Id @GeneratedValue(strategy = GenerationType.IDENTITY)
    Long id;
    String name;

    @ManyToMany(mappedBy = "interests")
    private Set<User> users;

    @ManyToMany(mappedBy = "interests")
    private Set<Post> posts;
    public Interest() {}
}
