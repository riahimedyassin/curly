:smile Template parser (Figure out if you are going to pass a config resolver or use the config values)
[] Parse templates and verify the template structure (tmpl file)
[] Resolve component final path (if not present with the user args and does not interfer with the config principle generate from the config).
[] Work on valiation. The validation should consider the values that may appear to be legit but interfer with the global end config. Verify that so that the services or the end consumers does not care about the validity of the data.