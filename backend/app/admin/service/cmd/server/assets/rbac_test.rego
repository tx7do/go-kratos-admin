package authz.introspection_test

import data.authz.introspection

import rego.v1

# 测试：授权成功
test_authorized_success if {
    mock_policies := {
        "user1": [
            {"pattern": "resource1", "method": "GET"}
        ]
    }

    test_input := {
        "subjects": ["user1"],
        "pairs": [
            {"resource": "resource1", "action": "GET"}
        ]
    }

    introspection.authorized with data.policies as mock_policies with input as test_input
}

# 测试：授权失败（资源不匹配）
test_authorized_resource_mismatch if {
    policies := {
        "user1": [
            {"pattern": "resource1", "method": "GET"}
        ]
    }

    test_input := {
        "subjects": ["user1"],
        "pairs": [
            {"resource": "resource2", "action": "GET"}
        ]
    }

    not data.authz.introspection.authorized with data.policies as policies with input as test_input
}

# 测试：授权失败（方法不匹配）
test_authorized_method_mismatch if {
    policies := {
        "user1": [
            {"pattern": "resource1", "method": "GET"}
        ]
    }

    test_input := {
        "subjects": ["user1"],
        "pairs": [
            {"resource": "resource1", "action": "POST"}
        ]
    }

    not data.authz.introspection.authorized with data.policies as policies with input as test_input
}

# 测试：授权项目
test_authorized_project if {
    policies := {
        "user1": [
            {"pattern": "resource1", "method": "GET"}
        ]
    }

    test_input := {
        "subjects": ["user1"],
        "pairs": [
            {"resource": "resource1", "action": "GET"}
        ]
    }

    data.authz.introspection.authorized_project with data.policies as policies with input as test_input == "api"
}

# 测试：授权对
test_authorized_pair if {
    policies := {
        "user1": [
            {"pattern": "resource1", "method": "GET"}
        ]
    }

    test_input := {
        "subjects": ["user1"],
        "pairs": [
            {"resource": "resource1", "action": "GET"}
        ]
    }

    data.authz.introspection.authorized_pair with data.policies as policies with input as test_input == [{"resource": "resource1", "action": "GET"}]
}