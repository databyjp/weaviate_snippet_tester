import assert from 'node:assert';
import weaviate from 'weaviate-ts-client';

// // For local client
// const client = weaviate.client({
//   scheme: 'http',
//   host: 'localhost:8080',
//   headers: {
//     'X-OpenAI-Api-Key': process.env['OPENAI_APIKEY'],
//     'X-Cohere-Api-Key': process.env['COHERE_APIKEY'],
//     'X-HuggingFace-Api-Key': process.env['HUGGINGFACE_APIKEY'],
//   },
// });

// // For edu-demo client
const client = weaviate.client({
  scheme: 'https',
  host: 'edu-demo.weaviate.network',
  apiKey: new weaviate.ApiKey('learn-weaviate'),
  headers: {
    'X-OpenAI-Api-Key': process.env['OPENAI_APIKEY'],
    'X-Cohere-Api-Key': process.env['COHERE_APIKEY'],
    'X-HuggingFace-Api-Key': process.env['HUGGINGFACE_APIKEY'],
  },
});

let result;

// ==============================
// ===== BASIC GET EXAMPLES =====
// ==============================

// BasicGetJS
result = await client
  .graphql
  .get()
  .withClassName('JeopardyQuestion')
  .withFields('question')
  .withLimit(2)
  .do();

console.log(JSON.stringify(result, null, 2));
// END BasicGetJS

// Test
assert('JeopardyQuestion' in result.data.Get);
let questionKeys = new Set(Object.keys(result.data.Get.JeopardyQuestion[0]));
assert.deepEqual(questionKeys, new Set(['question']));
// End test
