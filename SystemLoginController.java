package com.scity.pms.system.authenticate.interfaces.web;

import com.scity.pms.shared.dto.BasicResult;
import com.scity.pms.shared.dto.ResultCodeEnum;
import com.scity.pms.system.authenticate.domain.SysUserResource;
import com.scity.pms.system.authenticate.infrastructure.TokenManager;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import java.util.UUID;

import static org.springframework.http.HttpHeaders.AUTHORIZATION;

/**
 * {@code SystemLoginController} 包含了业务交换平台一期自身的简单登录登出
 *
 * @author hege
 * @version v1.0
 * @since v1.0
 */
@RestController
public class SystemLoginController {

    /**
     * 用户密码登录请求。
     * 目前用户密码是在配置文件写死
     *
     * @param command 主要包括用户名、md加密之后的密码等参数
     * @return 登录成功返回success
     */
    @PostMapping("/login")
    public BasicResult passwordLogin(@RequestBody @Validated PasswordLoginCommand command) {
        boolean checkPswd = command.getAccount().equals(adminAccount) && command.getEncryptedPassword().equals(encryptedPassword);
        if (!checkPswd) {
            return BasicResult.failWith(ResultCodeEnum.INCORRECT_CREDENTIAL);
        }
        String token = UUID.randomUUID().toString();
        tokenManager.cacheUserResource(token, new SimpleUserResource(1, "-", "超级管理员"));
        return BasicResult.successWith(new SimpleLoginResult(token));
    }

    /**
     * 登出。
     * 简单从 redis 删除登录信息
     *
     * @return always success
     */
    @PostMapping("/logout")
    public BasicResult logout(@RequestHeader(AUTHORIZATION) String token) {
        redisTemplate.delete(token);
        return BasicResult.success();
    }

    @Value("${system.admin.account:exchangeadmin01}")
    private String adminAccount;

    @Value("${system.admin.pswd:18eef1c32ec6871e1be82e3d7e268a82}")
    private String encryptedPassword;

    @Resource
    private RedisTemplate<String, Object> redisTemplate;

    @Resource
    private TokenManager tokenManager;

    public static class SimpleLoginResult {
        private final String token;

        public SimpleLoginResult(String token) {
            this.token = token;
        }

        public String getToken() {
            return token;
        }
    }

    public static class SimpleUserResource implements SysUserResource {

        private long userId ;

        private String deptCode ;

        private String username ;

        public SimpleUserResource() {
        }

        public SimpleUserResource(long userId, String deptCode, String username) {
            this.userId = userId;
            this.deptCode = deptCode;
            this.username = username;
        }

        public void setUserId(long userId) {
            this.userId = userId;
        }

        public void setDeptCode(String deptCode) {
            this.deptCode = deptCode;
        }

        public void setUsername(String username) {
            this.username = username;
        }

        @Override
        public long getUserId() {
            return userId;
        }

        @Override
        public String getDeptCode() {
            return deptCode;
        }

        @Override
        public String getUsername() {
            return username;
        }
    }
}
