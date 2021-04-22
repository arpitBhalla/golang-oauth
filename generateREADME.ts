const fs = require("fs");

type ReadmeContent = {
  name: string;
  endpoint: string;
  method: "POST" | "GET";
  request: {};
  response: {};
};

const content: ReadmeContent[] = [
  {
    name: "Register",
    endpoint: "register",
    method: "POST",
    request: {
      name: "YOUR NAME",
      email: "EMAIL",
      password: "PASSWORD",
    },
    response: {
      code: 200,
    },
  },
  {
    name: "Login",
    endpoint: "login",
    method: "POST",
    request: {
      email: "EMAIL",
      password: "PASSWORD",
    },
    response: {
      code: 200,
      accessToken: "JWT TOKEN",
      refreshToken: "JWT TOKEN",
    },
  },
  {
    name: "Refresh Token",
    endpoint: "refresh",
    method: "POST",
    request: {
      refreshToken: "JWT TOKEN",
    },
    response: {
      accessToken: "JWT TOKEN",
    },
  },
  {
    name: "Logout",
    endpoint: "logout",
    method: "POST",
    request: null,
    response: {
      code: 200,
    },
  },
  {
    name: "Get your profile",
    endpoint: "profile",
    method: "GET",
    request: null,
    response: {
      code: 200,
      _id: "MONGODB ID",
      name: "NAME",
      email: "EMAIL",
    },
  },
  {
    name: "Dump all profile",
    endpoint: "all",
    method: "GET",
    request: null,
    response: {
      code: 200,
      data: [
        {
          _id: "MONGODB ID",
          name: "NAME",
          email: "EMAIL",
        },
      ],
    },
  },
];

let md = `# API

## Base URL

> https://nitkkr-online.el.r.appspot.com/

## Authentication Header

 \`Bearer: <Token>\`

## Endpoints

`;

content.forEach((e) => {
  md += `
#### ${e.name}

> \`${e.method}\`  **/${e.endpoint}**

${
  e.request
    ? `**Request Body**

\`\`\`json
${JSON.stringify(e.request, null, 2)}
\`\`\``
    : ""
}

**Response**

\`\`\`json
${JSON.stringify(e.response, null, 2)}
\`\`\`

`;
});

fs.writeFileSync("README.md", md);
