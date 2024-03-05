import Configuration from '../config/config'
import React, { useEffect, useState } from "react";

const App = () => {
  const [data, setData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
      const fetchData = async () => {
        try {
          const response = await axios.get(`http://localhost:${Configuration.port}/api/webhook-2`);
          setData(response.data);
        } catch (error) {
          setError(error);
        } finally {
          setLoading(false);
        }
      };

      fetchData();
    }, []);

  const renderForm = () => {
    if (!Configuration || loading) {
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
          } else {
            return (
              <p key={name}>
                <b>{name}:</b> {webhook[name].toString()}
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
