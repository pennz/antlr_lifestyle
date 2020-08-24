FROM neelkamath/spacy-server:2-en_core_web_sm

ARG PORT
WORKDIR /app
ENV PYTHONBUFFERED=1 SENSE2VEC=0 SPACY_MODEL=en_core_web_sm

USER user
CMD ["sh", "-c", "uvicorn src.main:app --host 0.0.0.0 --port $PORT"]
