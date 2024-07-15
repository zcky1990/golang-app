import Dropdown from "../../static/javascript/form/Dropdown.vue";

export default {
  title: "Form/Dropdown",
  component: Dropdown,
  tags: ["autodocs"],
  parameters: {
    docs: {
      description: {
        component: "An Dropdown Component Example.",
      },
    },
  },
  render: (args) => ({
    components: { Dropdown },
    setup() {
      return { ...args };
    },
    template: `
      <div class="container">
        <Dropdown v-model:value="value" :dropdown-label="label" :dropdown-items="items" v-model:show-error="error" />
      </div>
    `
  }),
  argTypes: {
    value: { control: "text" },
    label: { control: "text" },
    items: { control: "object" },
    error: { 
      control: { type: 'boolean' }, 
    }
  },
};

export const DropdownWithLabel = {
  parameters: {
    docs: {
      description: {
        story: "A Dropdown component example with label",
      },
      source: {
        code: `<Dropdown v-model:value="value" :dropdown-label="label" :dropdown-items="items" v-model:show-error="error" />`,
      },
    },
  },
  args: {
    label: 'Example Dropdown Label',
    value: '',
    error: false,
    items: [
      {
        key: "Dropdown 1",
        value: "1",
      },
      {
        key: "Dropdown 2",
        value: "2",
      },
    ],
  },
};

export const DropdownWithoutLabel = {
  parameters: {
    docs: {
      description: {
        story: "A Dropdown component example without label",
      },
      source: {
        code: `<Dropdown v-model:value="value" :dropdown-items="items" v-model:show-error="error" />`,
      },
    },
  },
  args: {
    label: '',
    value: '',
    error: false,
    items: [
      {
        key: "Dropdown 1",
        value: "1",
      },
      {
        key: "Dropdown 2",
        value: "2",
      },
    ],
  },
};

export const DropdownWithLabelAndErrorShow = {
  parameters: {
    docs: {
      description: {
        story: "A Dropdown component example with label and error shown because no item selected",
      },
      source: {
        code: `<Dropdown v-model:value="value" :dropdown-label="label" :dropdown-items="items" v-model:show-error="error" />`,
      },
    },
  },
  args: {
    label: "Example Dropdown Label",
    value: '',
    error: true,
    items: [
      {
        key: "Dropdown 1",
        value: "1",
      },
      {
        key: "Dropdown 2",
        value: "2",
      },
    ],
  },
};

export const DropdownWithoutLabelAndErrorShow = {
  parameters: {
    docs: {
      description: {
        story: "A Dropdown component example without label and error shown because no item selected",
      },
      source: {
        code: `<Dropdown v-model:value="value" :dropdown-items="items" v-model:show-error="error" />`,
      },
    },
  },
  args: {
    label: '',
    value: '',
    error: true,
    items: [
      {
        key: "Dropdown 1",
        value: "1",
      },
      {
        key: "Dropdown 2",
        value: "2",
      },
    ],
  },
};
