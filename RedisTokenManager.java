package com.scity.pms.auth.infrastructure.token;

import com.scity.pms.auth.domain.login.LoginContext;
import com.scity.pms.auth.domain.login.LoginedUserResource;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.concurrent.TimeUnit;

/**
 * {@code RedisTokenManager}
 * Token manager, base on redis.
 *
 * @author Jackie Leung
 * @version 1.0
 * @since 1.0
 */
@Service("redisTokenManager")
public class RedisTokenManager implements TokenManager {

    @Resource
    private RedisTemplate<String, LoginedUserResource> redisTemplate;

    @Override
    public void put(LoginContext loginContext) {
        redisTemplate.opsForValue().set(
                loginContext.getKey(),
                loginContext.getUserSource(),
                loginContext.getCacheTime(),
                TimeUnit.SECONDS
        );
    }

    @Override
    public void remove(String token) {
        redisTemplate.delete(token);
    }

    @Override
    public boolean verifyToken(String token) {
        return redisTemplate.hasKey(token);
    }

    @Override
    public LoginedUserResource userSources(String token) {
        return redisTemplate.opsForValue().get(token);
    }
}
