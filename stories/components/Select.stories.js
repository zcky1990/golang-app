import SelectComponents from "../../static/javascript/components/Select.vue";

export default {
  title: "Components/Select",
  component: SelectComponents,
  tags: ["autodocs"],
  parameters: {
    docs: {
      description: {
        component: "An Select Component Example.",
      },
    },
  },
  render: (args) => ({
    components: { SelectComponents },
    setup() {
      return { ...args };
    },
    template: `
      <div class="container">
        <SelectComponents v-model:value="value" :select-label="label" :select-items="items" v-model:show-error="error" />
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

export const SelectWithLabel = {
  parameters: {
    docs: {
      description: {
        story: "A Select component example with label",
      },
      source: {
        code: `<SelectComponents v-model:value="value" :select-label="label" :select-items="items" v-model:show-error="error" />`,
      },
    },
  },
  args: {
    label: 'Example Select Label',
    value: '',
    error: false,
    items: [
      {
        key: "Select 1",
        value: "1",
      },
      {
        key: "Select 2",
        value: "2",
      },
    ],
  },
};

export const SelectWithoutLabel = {
  parameters: {
    docs: {
      description: {
        story: "A Select component example without label",
      },
      source: {
        code: `<SelectComponents v-model:value="value" :select-items="items" v-model:show-error="error" />`,
      },
    },
  },
  args: {
    label: '',
    value: '',
    error: false,
    items: [
      {
        key: "Select 1",
        value: "1",
      },
      {
        key: "Select 2",
        value: "2",
      },
    ],
  },
};

export const SelectWithLabelAndErrorShow = {
  parameters: {
    docs: {
      description: {
        story: "A Select component example with label and error shown because no item selected",
      },
      source: {
        code: `<SelectComponents v-model:value="value" :select-label="label" :select-items="items" v-model:show-error="error" />`,
      },
    },
  },
  args: {
    label: "Example Select Label",
    value: '',
    error: true,
    items: [
      {
        key: "Select 1",
        value: "1",
      },
      {
        key: "Select 2",
        value: "2",
      },
    ],
  },
};

export const SelectWithoutLabelAndErrorShow = {
  parameters: {
    docs: {
      description: {
        story: "A Select component example without label and error shown because no item selected",
      },
      source: {
        code: `<SelectComponents v-model:value="value" :select-items="items" v-model:show-error="error" />`,
      },
    },
  },
  args: {
    label: '',
    value: '',
    error: true,
    items: [
      {
        key: "Select 1",
        value: "1",
      },
      {
        key: "Select 2",
        value: "2",
      },
    ],
  },
};
