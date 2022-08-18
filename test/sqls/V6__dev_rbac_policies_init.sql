INSERT INTO
    idk_rbac_policies (ptype, v0, v1, v2, v3, v4, v5)
VALUES
    ("p", "anonymous", "/client/.*", "GET", "allow", "", "");

INSERT INTO
    idk_rbac_policies (ptype, v0, v1, v2, v3, v4, v5)
VALUES
    ("p", "anonymous", "/client/.*/metrices(\\/[1].*)*", "GET", "deny", "", "");

INSERT INTO
    idk_rbac_policies (ptype, v0, v1, v2, v3, v4, v5)
VALUES
    ("p", "anonymous", "/client/.*/comment", "(POST)|(DELETE)", "allow", "", "");

INSERT INTO
    idk_rbac_policies (ptype, v0, v1, v2, v3, v4, v5)
VALUES
    ("p", "anonymous", "/client/.*/comment", "(PUT)|(PATCH)", "deny", "", "");

INSERT INTO
    idk_rbac_policies (ptype, v0, v1, v2, v3, v4, v5)
VALUES
    ("p", "anonymous", "/client/.*/ws", "(POST)|(GET)", "allow", "", "");