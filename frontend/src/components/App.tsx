import Configuration from '../config/config'

const App = () => {
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
          } else if (name == "some_thing_else") {
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
