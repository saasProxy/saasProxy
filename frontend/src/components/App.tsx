import Configuration from '../config/config'
import React, {useEffect, useState} from "react";
import axios from 'axios';

const App = () => {
  // const [data, setData] = useState(null);
  // const [loading, setLoading] = useState(true);
  // const [error, setError] = useState(null);

  const handleClick = async (verb: string, slug: string) => {
    try {
      const response = async () => {
        if (verb == "POST") {
          return axios.post(`http://localhost:${Configuration.port}${slug}`.toString(), {});
        } else {
          return axios.get(`http://localhost:${Configuration.port}${slug}`.toString());
        }
      }

      const result = await response();

      // Handle the fetched data or perform other async tasks
      console.log(`[${slug}] async operation result`, result);
      return result;
    } catch (error) {
      console.error(`[${slug}] error during async operation`, error);
    }
  };

  const renderForm = () => {
    if (!Configuration) {
      return <div>Loading...</div>;
    }

    let webhookKeys = Object.keys(Configuration.webhooks[0]);

    return Configuration.webhooks.map((webhook) => (
      <div>
        <h1>Webhook Configuration</h1>
        {webhookKeys.map((name) => {
          if (name == "headers") {
            return (
              <p key={name}>
                <b>{name}:</b> {JSON.stringify(webhook[name], null, 2)}
              </p>
            )
          } else if (name == "incoming_slug") {
            return (
              <p key={name}>
                <b>{name}:</b> {webhook[name].toString()}
              </p>
            )
          } else if (name == "request_verb") {
            return (
              <p key={name}>
                <b>{name}: </b><button onClick={() => handleClick(webhook.request_verb, webhook.incoming_slug)}>{webhook.request_verb}</button>
              </p>
            )
          } else {
            return (
              <p key={name}>
                <b>{name}:</b> {
                // @ts-expect-error
                webhook[name].toString()
              }
              </p>
            )
          }
        })}
        <hr/>
      </div>
    ));
  };

  return (
    <div className="">
      <div className="">
        <p><b>Port:</b> {Configuration.port}</p>
        <p><b>Destination:</b> {Configuration.destination}</p>
        <hr/>
        {renderForm()}
      </div>
    </div>
  );
}

export default App
