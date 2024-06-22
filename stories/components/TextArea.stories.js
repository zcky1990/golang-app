import TextArea from "../../static/javascript/components/TextArea.vue";

export default {
  title: "Components/TextArea",
  component: TextArea,
  tags: ["autodocs"],
  parameters: {
    docs: {
      description: {
        component: "An TextArea Component Example. This component can be used to capture user TextArea and handle validation errors.",
      },
    },
  },
  render: (args) => ({
    components: {
      TextArea,
    },
    setup() {
      return {
        ...args,
      };
    },
    template:
      '<TextArea v-model:value="value" :text-area-label="label" v-model:show-error="error" :text-area-error-label="labelError" />',
  }),
  argTypes: {
    value: { control: "text" },
    label: { control: "text" },
    error: {
      options: [true, false],
      control: { type: "select" },
    },
    labelError: { control: "text" },
  },
};

export const TextAreaWithLabel = {
  parameters: {
    docs: {
      description: {
        story: "An Example Text Area component with label.",
      },
      source: {
        code: `<TextArea v-model:value="value" :text-area-label="label" v-model:show-error="error" :text-area-error-label="labelError" />`
      },
    },
  },
  args: {
    value: "Example text area value",
    label: "Text Area label Example",
    error: false,
    labelError: "Example error message",
  },
};

export const TextAreaWithoutLabel = {
  parameters: {
    docs: {
      description: {
        story: "An Example Text Area component without label",
      },
      source: {
        code: `<TextArea v-model:value="value" :text-area-label="label" v-model:show-error="error" :text-area-error-label="labelError" />`
      },
    },
  },
  args: {
    value: "Example text area value",
    label: "",
    error: false,
    labelError: "Example error message",
  },
};

export const TextAreaWithLabelAndError = {
  parameters: {
    docs: {
      description: {
        story: "An Example Text Area component with label and Error.",
      },
      source: {
        code: `<TextArea v-model:value="value" :text-area-label="label" v-model:show-error="error" :text-area-error-label="labelError" />`
      },
    },
  },
  args: {
    value: "Example text area value",
    label: "Text Area label Example",
    error: true,
    labelError: "Example error message",
  },
};

export const TextAreaWithoutLabelAndError = {
  parameters: {
    docs: {
      description: {
        story: "An Example Text Area component without label",
      },
      source: {
        code: `<TextArea v-model:value="value" :text-area-label="label" v-model:show-error="error" :text-area-error-label="labelError" />`
      },
    },
  },
  args: {
    value: "Example text area value",
    label: "",
    error: true,
    labelError: "Example error message",
  },
};