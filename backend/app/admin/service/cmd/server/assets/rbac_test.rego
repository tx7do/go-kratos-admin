package authz.introspection

import future.keywords.if

# 测试：授权成功
test_authorized_success {
    data.policies := {
        "user1": [
            {"pattern": "resource1", "method": "read"}
        ]
    }

    input := {
        "subjects": ["user1"],
        "pairs": [
            {"resource": "resource1", "action": "read"}
        ]
    }

    authorized
}

# 测试：授权失败（资源不匹配）
test_authorized_resource_mismatch {
    data.policies := {
        "user1": [
            {"pattern": "resource1", "method": "read"}
        ]
    }

    input := {
        "subjects": ["user1"],
        "pairs": [
            {"resource": "resource2", "action": "read"}
        ]
    }

    not authorized
}

# 测试：授权失败（方法不匹配）
test_authorized_method_mismatch {
    data.policies := {
        "user1": [
            {"pattern": "resource1", "method": "read"}
        ]
    }

    input := {
        "subjects": ["user1"],
        "pairs": [
            {"resource": "resource1", "action": "write"}
        ]
    }

    not authorized
}

# 测试：授权项目
test_authorized_project {
    data.policies := {
        "user1": [
            {"pattern": "resource1", "method": "read"}
        ]
    }

    input := {
        "subjects": ["user1"],
        "pairs": [
            {"resource": "resource1", "action": "read"}
        ]
    }

    authorized_project == "api"
}

# 测试：授权对
test_authorized_pair {
    data.policies := {
        "user1": [
            {"pattern": "resource1", "method": "read"}
        ]
    }

    input := {
        "subjects": ["user1"],
        "pairs": [
            {"resource": "resource1", "action": "read"}
        ]
    }

    authorized_pair == [{"resource": "resource1", "action": "read"}]
}