# %%

import pandas as pd
import sqlalchemy

con = sqlalchemy.create_engine("sqlite:///database_test.db")
# %%

query = """
SELECT 
    UUID as uuid,
    null AS twitch_id,
    NickTwitch AS twitch_nick

FROM tb_users
"""

df = pd.read_sql(query, con)

# %%
df.to_sql("twitch_users", con, index=False, if_exists="append")
# %%
