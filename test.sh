#!/bin/bash

docker run -it --rm quickforge zsh -i -c "
  echo '\n🔧 Step 1: Initial run';
  /app/quickforge -y;

  echo '\n🔧 Step 2: Source shell and run again';
  source /root/.zshrc;
  /app/quickforge -y;

  echo '\n✅ E2E complete';
"