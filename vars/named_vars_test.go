package vars_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/concourse/concourse/vars"
)

var _ = Describe("NamedVariables", func() {
	Describe("Get", func() {
		It("return no value and not found if there are no sources", func() {
			val, found, err := NewNamedVariables(nil).Get(VariableDefinition{})
			Expect(val).To(BeNil())
			Expect(found).To(BeFalse())
			Expect(err).ToNot(HaveOccurred())
		})

		It("return no value and not found if var source name doesn's exist", func() {
			vars1 := StaticVariables{"key1": "val"}
			vars2 := StaticVariables{"key2": "val"}
			vars := NewNamedVariables(map[string]Variables{"s1": vars1, "s2": vars2})

			val, found, err := vars.Get(VariableDefinition{Name: "s3:key1"})
			Expect(val).To(BeNil())
			Expect(found).To(BeFalse())
			Expect(err).ToNot(HaveOccurred())
		})

		It("return no value and not found if var source name is not specified", func() {
			vars1 := StaticVariables{"key1": "val"}
			vars2 := StaticVariables{"key2": "val"}
			vars := NewNamedVariables(map[string]Variables{"s1": vars1, "s2": vars2})

			val, found, err := vars.Get(VariableDefinition{Name: "key1"})
			Expect(val).To(BeNil())
			Expect(found).To(BeFalse())
			Expect(err).ToNot(HaveOccurred())
		})

		It("return error as soon as one source fails", func() {
			vars1 := StaticVariables{"key1": "val"}
			vars2 := &FakeVariables{GetErr: errors.New("fake-err")}
			vars := NewNamedVariables(map[string]Variables{"s1": vars1, "s2": vars2})

			val, found, err := vars.Get(VariableDefinition{Name: "s2:key3"})
			Expect(val).To(BeNil())
			Expect(found).To(BeFalse())
			Expect(err).To(Equal(errors.New("fake-err")))
		})

		It("return found value as soon as one source succeeds", func() {
			vars1 := &FakeVariables{}
			vars2 := StaticVariables{"key2": "val"}
			vars3 := &FakeVariables{GetErr: errors.New("fake-err")}
			vars := NewNamedVariables(map[string]Variables{"s1": vars1, "s2": vars2, "s3": vars3})

			val, found, err := vars.Get(VariableDefinition{Name: "s2:key2"})
			Expect(val).To(Equal("val"))
			Expect(found).To(BeTrue())
			Expect(err).ToNot(HaveOccurred())

			// Didn't get past other variables
			Expect(vars1.GetCallCount).To(Equal(0))
			Expect(vars3.GetCallCount).To(Equal(0))
		})
	})

	Describe("List", func() {
		It("returns list of names from multiple vars with duplicates", func() {
			defs, err := NewNamedVariables(nil).List()
			Expect(defs).To(BeEmpty())
			Expect(err).ToNot(HaveOccurred())

			vars := NewNamedVariables(map[string]Variables{
				"s1": StaticVariables{"a": "1", "b": "2"},
				"s2": StaticVariables{"b": "3", "c": "4"},
			})

			defs, err = vars.List()
			Expect(defs).To(ConsistOf([]VariableDefinition{{Name: "a"}, {Name: "b"}, {Name: "b"}, {Name: "c"}}))
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
