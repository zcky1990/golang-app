import Input from "../../static/javascript/components/Input.vue";

export default {
  title: "Components/Input",
  component: Input,
  tags: ["autodocs"],
  parameters: {
    docs: {
      description: {
        component: "An Input Component Example. This component can be used to capture user input and handle validation errors.",
      },
    },
  },
  render: (args) => ({
    components: {
      Input,
    },
    setup() {
      return {
        ...args,
      };
    },
    template:
      '<Input v-model:value="value" :input-label="label" :input-type="type" :show-error="showError" :input-error-label="messageError" />',
  }),
  argTypes: {
    value: { control: "text" },
    label: { control: "text" },
    type: {
      options: ["text", "password", "tel", "email", "number"],
      control: { type: "select" },
    },
    showError: {
      options: [true, false],
      control: { type: "select" },
    },
    messageError: { control: "text" },
  },
};

export const Text = {
  parameters: {
    docs: {
      description: {
        story: "An input component example with text as the input type.",
      },
      source: {
        code: `<Input v-model:value="value" input-label="Username" input-type="text" :show-error="false" input-error-label="Example error message" />`
      },
    },
  },
  args: {
    value: "Example input value",
    label: "Username",
    type: "text",
    showError: false,
    messageError: "Example error message",
  },
};

export const TextWithError = {
  parameters: {
    docs: {
      description: {
        story: "An input component example with text as the input type and showing an error message.",
      },
      source: {
        code: `<Input v-model:value="value" input-label="Username" input-type="text" :show-error="true" input-error-label="Example error message" />`
      },
    },
  },
  args: {
    value: "Example input value",
    label: "Username",
    type: "text",
    showError: true,
    messageError: "Example error message",
  },
};

export const Password = {
  parameters: {
    docs: {
      description: {
        story: "An input component example with password as the input type.",
      },
      source: {
        code: `<Input v-model:value="value" input-label="Password" input-type="password" :show-error="false" input-error-label="Example error message" />`
      },
    },
  },
  args: {
    value: "Example input value",
    label: "Password",
    type: "password",
    showError: false,
    messageError: "Example error message",
  },
};

export const PasswordWithError = {
  parameters: {
    docs: {
      description: {
        story: "An input component example with password as the input type and showing an error message.",
      },
      source: {
        code: `<Input v-model:value="value" input-label="Password" input-type="password" :show-error="true" input-error-label="Example error message" />`
      },
    },
  },
  args: {
    value: "Example input value",
    label: "Password",
    type: "password",
    showError: true,
    messageError: "Example error message",
  },
};

export const Email = {
  parameters: {
    docs: {
      description: {
        story: "An input component example with email as the input type.",
      },
      source: {
        code: `<Input v-model:value="value" input-label="Email" input-type="email" :show-error="false" input-error-label="Example error message" />`
      },
    },
  },
  args: {
    value: "example@mail.com",
    label: "Email",
    type: "email",
    showError: false,
    messageError: "Example error message",
  },
};

export const EmailWithError = {
  parameters: {
    docs: {
      description: {
        story: "An input component example with email as the input type and showing an error message.",
      },
      source: {
        code: `<Input v-model:value="value" input-label="Email" input-type="email" :show-error="true" input-error-label="Example error message" />`
      },
    },
  },
  args: {
    value: "example@mail.com",
    label: "Email",
    type: "email",
    showError: true,
    messageError: "Example error message",
  },
};
