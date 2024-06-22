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
      '<Input v-model:value="value" :input-label="label" :input-type="type" v-model:show-error="showError" :input-error-label="messageError" />',
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

export const InputWithTextAsType = {
  parameters: {
    docs: {
      description: {
        story: "An input component example with text as the input type.",
      },
      source: {
        code: `<Input v-model:value="value" input-label="Username" input-type="text" v-model:show-error="false" input-error-label="Example error message" />`
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

export const InputWithTextAsTypeWithError = {
  parameters: {
    docs: {
      description: {
        story: "An input component example with text as the input type and showing an error message.",
      },
      source: {
        code: `<Input v-model:value="value" input-label="Username" input-type="text" v-model:show-error="true" input-error-label="Example error message" />`
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

export const InputWithPasswordAsType = {
  parameters: {
    docs: {
      description: {
        story: "An input component example with password as the input type.",
      },
      source: {
        code: `<Input v-model:value="value" input-label="Password" input-type="password" v-model:show-error="false" input-error-label="Example error message" />`
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

export const InputWithPasswordAsTypeWithError = {
  parameters: {
    docs: {
      description: {
        story: "An input component example with password as the input type and showing an error message.",
      },
      source: {
        code: `<Input v-model:value="value" input-label="Password" input-type="password" v-model:show-error="true" input-error-label="Example error message" />`
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

export const InputWithEmailAsType = {
  parameters: {
    docs: {
      description: {
        story: "An input component example with email as the input type.",
      },
      source: {
        code: `<Input v-model:value="value" input-label="Email" input-type="email" v-model:show-error="false" input-error-label="Example error message" />`
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

export const InputWithEmailAsTypeWithError = {
  parameters: {
    docs: {
      description: {
        story: "An input component example with email as the input type and showing an error message.",
      },
      source: {
        code: `<Input v-model:value="value" input-label="Email" input-type="email" v-model:show-error="true" input-error-label="Example error message" />`
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
