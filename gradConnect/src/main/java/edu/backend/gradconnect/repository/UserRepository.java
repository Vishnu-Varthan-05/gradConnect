package edu.backend.gradconnect.repository;

import edu.backend.gradconnect.models.User;
import org.springframework.data.repository.CrudRepository;


public interface UserRepository extends CrudRepository<User, Integer> {
    User getUserByEmail(String email);
}
