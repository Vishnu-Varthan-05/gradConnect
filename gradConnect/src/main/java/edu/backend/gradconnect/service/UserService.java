package edu.backend.gradconnect.service;

import edu.backend.gradconnect.models.User;
import edu.backend.gradconnect.repository.UserRepository;
import org.springframework.stereotype.Service;

@Service
public class UserService {
    private final UserRepository userRepository;
    public UserService(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

}
